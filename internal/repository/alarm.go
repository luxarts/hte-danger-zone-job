package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/domain"
	"log"
	"time"
)

type AlarmRepository interface {
	Send(deviceID string, companyID string, message string) error
}

type alarmRepository struct {
	rc          *redis.Client
	alarmsQueue string
}

func NewAlarmRepository(rc *redis.Client, alarmsQueue string) AlarmRepository {
	return &alarmRepository{
		rc:          rc,
		alarmsQueue: alarmsQueue,
	}
}

func (repo *alarmRepository) Send(deviceID string, companyID string, message string) error {
	log.Printf("%s->%s\n", deviceID, message)

	ctx := context.Background()

	alarm := domain.Alarm{
		AssetID:   deviceID,
		Type:      "dangerzone",
		Action:    "create",
		Timestamp: time.Now().Unix(),
		CompanyID: companyID,
		Position:  nil,
		Text:      message,
		CountryID: 0,
		Device:    "",
	}

	alarmBytes, err := json.Marshal(alarm)
	if err != nil {
		return err
	}

	err = repo.rc.RPush(ctx, repo.alarmsQueue, string(alarmBytes)).Err()

	return err
}
