package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
)

type DangerZoneRepository interface {
	GetByUserID(userID string) (*domain.DangerZone, error)
}

type dangerZoneRepository struct {
	rc *redis.Client
}

func NewDangerZoneRepository(rc *redis.Client) DangerZoneRepository {
	return &dangerZoneRepository{rc: rc}
}

func (repo *dangerZoneRepository) GetByUserID(userID string) (*domain.DangerZone, error) {
	ctx := context.Background()
	res, err := repo.rc.
		Get(ctx, fmt.Sprintf("%s:%s", defines.DangerZoneKey, userID)).
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
