package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
)

type DangerZoneRepository interface {
	GetAllActive() (*[]domain.DangerZone, error)
}

type dangerZoneRepository struct {
	baseURL string
	rc      *resty.Client
}

func NewDangerZoneRepository(rc *resty.Client, baseURL string) DangerZoneRepository {
	return &dangerZoneRepository{baseURL: baseURL, rc: rc}
}
func (repo *dangerZoneRepository) GetAllActive() (*[]domain.DangerZone, error) {
	resp, err := repo.rc.R().Get(fmt.Sprintf("%s%s", repo.baseURL, defines.APIDangerZoneGetAll))
	if err != nil {
		return nil, err
	}

	var dzs []domain.DangerZone
	err = json.Unmarshal(resp.Body(), &dzs)
	if err != nil {
		return nil, err
	}

	return &dzs, nil
}
