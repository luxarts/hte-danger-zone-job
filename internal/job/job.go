package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/controller"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
	"hte-danger-zone-job/internal/service"
	"log"
	"os"
	"sync"
	"time"
)

type Job interface {
	Run()
}

type job struct {
	redisClient            *redis.Client
	redisChannelGeoloc     string
	redisChannelNewZone    string
	redisChannelDeleteZone string
	gCtrl                  controller.GeolocController
	dzcCtrl                controller.DangerZoneCacheController
	dzCtrl                 controller.DangerZoneController
	alarmCtrl              controller.AlarmController
	cancelMap              map[string]context.CancelFunc
	cancelMapMutex         sync.Mutex
}

func New() Job {
	// Redis Client
	redisClient := initRedisClient()
	restClient := resty.New()

	// Repository
	alarmRepo := repository.NewAlarmRepository(redisClient, os.Getenv(defines.EnvRedisQueueAlarms))
	dzcRepo := repository.NewDangerZoneCacheRepository(redisClient, os.Getenv(defines.EnvRedisKeyDangerZone))
	dzRepo := repository.NewDangerZoneRepository(restClient, os.Getenv(defines.EnvAPIDangerZoneHost))

	// Service
	zoneSvc := service.NewZoneService(dzcRepo, alarmRepo)
	dzcSvc := service.NewDangerZoneCacheService(dzcRepo)
	dzSvc := service.NewDangerZoneService(dzRepo)
	alarmSvc := service.NewAlarmService(alarmRepo)

	// Controller
	gCtrl := controller.NewGeolocController(zoneSvc)
	dzcCtrl := controller.NewDangerZoneCacheController(dzcSvc)
	dzCtrl := controller.NewDangerZoneController(dzSvc)
	alarmCtrl := controller.NewAlarmController(alarmSvc)

	return &job{
		gCtrl:                  gCtrl,
		dzcCtrl:                dzcCtrl,
		dzCtrl:                 dzCtrl,
		alarmCtrl:              alarmCtrl,
		redisClient:            redisClient,
		redisChannelGeoloc:     os.Getenv(defines.EnvRedisChannelGeoloc),
		redisChannelNewZone:    os.Getenv(defines.EnvRedisChannelCreateDangerZone),
		redisChannelDeleteZone: os.Getenv(defines.EnvRedisChannelDeleteDangerZone),
		cancelMap:              make(map[string]context.CancelFunc),
	}
}

func (j *job) Run() {
	go j.recreateActiveZonesJob()
	go j.createZoneJob()
	go j.deleteZoneJob()

	select {}
}

func (j *job) recreateActiveZonesJob() {
	log.Printf("Recreating active zones\n")
	dangerZones, err := j.dzCtrl.GetAllActive()
	if err != nil {
		log.Panicf("Failed GetAllActive: %+v\n", err)
	}
	for _, dz := range *dangerZones {
		// Store zone in cache
		err = j.dzcCtrl.Create(&dz)
		if err != nil {
			log.Println(err)
			continue
		}
		ts := time.Unix(dz.EndTimestamp, 0)
		go j.createConsumerForDevice(dz.DeviceID, ts)
	}
}
func (j *job) createZoneJob() {
	log.Printf("Listening channel %s\n", j.redisChannelNewZone)
	for {
		ctx := context.Background()
		res, err := j.redisClient.Subscribe(ctx, j.redisChannelNewZone).ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		var body domain.DangerZone

		err = json.Unmarshal([]byte(res.Payload), &body)
		if err != nil {
			log.Println(err)
			continue
		}
		ts := time.Unix(body.EndTimestamp, 0)

		// Check if zone consumer already exists
		j.cancelMapMutex.Lock()
		_, exists := j.cancelMap[body.DeviceID]
		j.cancelMapMutex.Unlock()
		if exists {
			log.Println("Go routine already exists for this device")
			continue
		}

		// Store zone in cache
		err = j.dzcCtrl.Create(&body)
		if err != nil {
			log.Println(err)
			continue
		}

		go j.createConsumerForDevice(body.DeviceID, ts)
	}
}
func (j *job) deleteZoneJob() {
	log.Printf("Listening channel %s\n", j.redisChannelDeleteZone)
	for {
		ctx := context.Background()
		res, err := j.redisClient.Subscribe(ctx, j.redisChannelDeleteZone).ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}

		dz, err := j.dzcCtrl.GetByDeviceID(res.Payload)
		if err != nil {
			log.Println(err)
			continue
		}
		if dz == nil {
			continue
		}

		// Create go routine
		go j.deleteConsumerForDevice(dz.DeviceID)
	}
}

func initRedisClient() *redis.Client {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(defines.EnvRedisHost),
		Password: os.Getenv(defines.EnvRedisPassword),
	})

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Error ping Redis: %+v\n", err)
	}

	return redisClient
}
func (j *job) createConsumerForDevice(deviceID string, ts time.Time) {
	log.Printf("Created routine for deviceID: %s\n", deviceID)

	var ctx context.Context
	j.cancelMapMutex.Lock()
	ctx, j.cancelMap[deviceID] = context.WithDeadline(context.Background(), ts)
	j.cancelMapMutex.Unlock()

	msgChan := make(chan *redis.Message)

	ps := j.redisClient.Subscribe(ctx, fmt.Sprintf("%s:%s", j.redisChannelGeoloc, deviceID))

	go func() {
		for msg := range ps.Channel() {
			msgChan <- msg
		}
	}()

	for {
		select {
		case msg := <-msgChan:
			resp, err := j.gCtrl.Process(deviceID, msg.Payload)
			if err != nil {
				log.Printf("Error processing: %+v\n", err)
			}
			if resp != nil {
				err = j.dzCtrl.DeleteByDeviceID(deviceID)
				if err != nil {
					log.Printf("Error deleting from ms: %+v\n", err)
				}
				j.deleteConsumerForDevice(deviceID)
			}
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				err := j.alarmCtrl.Send(&domain.SendAlarmReq{
					DeviceID: deviceID,
					Message:  defines.AlarmMessageInsideZoneAfterTime,
				})
				if err != nil {
					log.Printf("Error sending alarm: %+v\n", err)
					return
				}

				err = j.dzCtrl.DeleteByDeviceID(deviceID)
				if err != nil {
					log.Printf("Error deleting from ms: %+v\n", err)
				}
			}
			err := ps.Close()
			if err != nil {
				log.Printf("Error closing PubSub: %+v\n", err)
			}
			log.Printf("Deleted routine for deviceID: %s\n", deviceID)
			return
		}
	}
}
func (j *job) deleteConsumerForDevice(id string) {
	j.cancelMapMutex.Lock()
	defer j.cancelMapMutex.Unlock()
	cancel, exist := j.cancelMap[id]
	if !exist {
		return
	}
	cancel()
	delete(j.cancelMap, id)

	err := j.dzcCtrl.DeleteByDeviceID(id)
	if err != nil {
		log.Printf("Error deleting from cache: %+v\n", err)
	}
}
