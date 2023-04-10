package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
)

type DangerZoneCacheRepository interface {
	Create(z *domain.DangerZone) error
	GetByDeviceID(deviceID string) (*domain.DangerZone, error)
}

type dangerZoneCacheRepository struct {
	key string
	rc  *redis.Client
}

func NewDangerZoneCacheRepository(rc *redis.Client, key string) DangerZoneCacheRepository {
	return &dangerZoneCacheRepository{rc: rc, key: key}
}

func (repo *dangerZoneCacheRepository) Create(z *domain.DangerZone) error {
	zBytes, err := json.Marshal(z)
	if err != nil {
		return err
	}

	ctx := context.Background()
	return repo.rc.Set(ctx, fmt.Sprintf("%s:%s", repo.key, z.DeviceID), zBytes, 0).Err()
}
func (repo *dangerZoneCacheRepository) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	ctx := context.Background()
	res, err := repo.rc.
		Get(ctx, fmt.Sprintf("%s:%s", defines.DangerZoneKey, deviceID)).
		Result()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	var dangerZone domain.DangerZone
	err = json.Unmarshal([]byte(res), &dangerZone)

	if err != nil {
		return nil, err
	}

	return &dangerZone, nil
}
