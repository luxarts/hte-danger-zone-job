package service

import (
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
)

type DangerZoneService interface {
	GetAllActive() (*[]domain.DangerZone, error)
}

type dangerZoneService struct {
	repo repository.DangerZoneRepository
}

func NewDangerZoneService(repo repository.DangerZoneRepository) DangerZoneService {
	return &dangerZoneService{repo: repo}
}
func (svc *dangerZoneService) GetAllActive() (*[]domain.DangerZone, error) {
	return svc.repo.GetAllActive()
}
