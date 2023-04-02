package job

import (
	"context"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/controller"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/repository"
	"hte-danger-zone-job/internal/service"
	"log"
	"os"
)

type Job interface {
	Run()
}

type job struct {
	redisClient *redis.Client
	redisStream string
	geolocCtrl  controller.GeolocController
}

func New() Job {
	// Redis Client
	redisClient := initRedisClient()

	// Repository
	zoneRepo := repository.NewDangerZoneRepository(redisClient)
	alarmRepo := repository.NewAlarmRepository()

	// Service
	zoneSvc := service.NewZoneService(zoneRepo, alarmRepo)

	// Controller
	ctrl := controller.NewGeolocController(zoneSvc)

	return &job{
		geolocCtrl:  ctrl,
		redisClient: redisClient,
		redisStream: os.Getenv(defines.EnvRedisStream),
	}
}

func (j *job) Run() {
	log.Printf("Listening stream %s\n", j.redisStream)
	for {
		ctx := context.Background()
		res, err := j.redisClient.XRead(ctx, &redis.XReadArgs{
			Streams: []string{j.redisStream, "$"},
			Block:   0,
		}).Result()

		if err != nil {
			log.Println(err)
			continue
		}

		if res == nil || len(res) != 1 {
			log.Printf("Invalid format: %+v\n", res)
			continue
		}

		for _, m := range res[0].Messages {
			go func(m redis.XMessage) {
				err := j.geolocCtrl.Process(&m)

				if err != nil {
					log.Printf("Error: %+v\n", err)
				}
			}(m)
		}

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
