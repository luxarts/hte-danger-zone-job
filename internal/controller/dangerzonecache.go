package controller

import (
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/service"
)

type DangerZoneCacheController interface {
	Create(body *domain.DangerZone) error
	GetByDeviceID(deviceID string) (*domain.DangerZone, error)
	DeleteByDeviceID(deviceID string) error
}
type dangerZoneCacheController struct {
	dzcSvc service.DangerZoneCacheService
}

func NewDangerZoneCacheController(dzcSvc service.DangerZoneCacheService) DangerZoneCacheController {
	return &dangerZoneCacheController{dzcSvc: dzcSvc}
}

func (ctrl *dangerZoneCacheController) Create(body *domain.DangerZone) error {
	return ctrl.dzcSvc.Create(body)
}
func (ctrl *dangerZoneCacheController) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	return ctrl.dzcSvc.GetByDeviceID(deviceID)
}
func (ctrl *dangerZoneCacheController) DeleteByDeviceID(deviceID string) error {
	return ctrl.dzcSvc.DeleteByDeviceID(deviceID)
}
