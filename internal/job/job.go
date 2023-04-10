package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/controller"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
	"hte-danger-zone-job/internal/service"
	"log"
	"os"
)

type Job interface {
	Run()
}

type job struct {
	redisClient         *redis.Client
	redisChannelGeoloc  string
	redisChannelNewZone string
	gCtrl               controller.GeolocController
	dzcCtrl             controller.DangerZoneCacheController
}

func New() Job {
	// Redis Client
	redisClient := initRedisClient()

	// Repository
	alarmRepo := repository.NewAlarmRepository()
	dzcRepo := repository.NewDangerZoneCacheRepository(redisClient, os.Getenv(defines.EnvRedisKeyDangerZone))

	// Service
	zoneSvc := service.NewZoneService(dzcRepo, alarmRepo)
	dzcSvc := service.NewDangerZoneCacheService(dzcRepo)

	// Controller
	gCtrl := controller.NewGeolocController(zoneSvc)
	dzcCtrl := controller.NewDangerZoneCacheController(dzcSvc)

	return &job{
		gCtrl:               gCtrl,
		dzcCtrl:             dzcCtrl,
		redisClient:         redisClient,
		redisChannelGeoloc:  os.Getenv(defines.EnvRedisChannelGeoloc),
		redisChannelNewZone: os.Getenv(defines.EnvRedisChannelCreateDangerZone),
	}
}

func (j *job) Run() {
	log.Printf("Listening channel %s\n", j.redisChannelNewZone)

	// Check new zones to be created
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

		// Check if zone already cached
		dz, err := j.dzcCtrl.GetByDeviceID(body.DeviceID)
		if err != nil {
			log.Println(err)
			continue
		}
		if dz != nil {
			continue
		}

		// Store zone in cache
		err = j.dzcCtrl.Create(&body)
		if err != nil {
			log.Println(err)
			continue
		}

		// Create go routine
		go j.createConsumerForDevice(body.DeviceID)
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
func (j *job) createConsumerForDevice(deviceID string) {
	log.Printf("Created routine for deviceID: %s\n", deviceID)
	for {
		ctx := context.Background()
		res, err := j.redisClient.
			Subscribe(ctx, fmt.Sprintf("%s:%s", j.redisChannelGeoloc, deviceID)).
			ReceiveMessage(ctx)
		if err != nil {
			log.Printf("Error subscribing: %+v\n", err)
			continue
		}

		err = j.gCtrl.Process(deviceID, res.Payload)
		if err != nil {
			log.Printf("Error processing: %+v\n", err)
		}
	}
}
