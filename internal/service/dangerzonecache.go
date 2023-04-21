package service

import (
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
)

type DangerZoneCacheService interface {
	Create(z *domain.DangerZone) error
	GetByDeviceID(deviceID string) (*domain.DangerZone, error)
	DeleteByDeviceID(deviceID string) error
}
type dangerZoneCacheService struct {
	repo repository.DangerZoneCacheRepository
}

func NewDangerZoneCacheService(repo repository.DangerZoneCacheRepository) DangerZoneCacheService {
	return &dangerZoneCacheService{repo: repo}
}
func (svc *dangerZoneCacheService) Create(z *domain.DangerZone) error {
	return svc.repo.Create(z)
}
func (svc *dangerZoneCacheService) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	return svc.repo.GetByDeviceID(deviceID)
}
func (svc *dangerZoneCacheService) DeleteByDeviceID(deviceID string) error {
	return svc.repo.DeleteByDeviceID(deviceID)
}
