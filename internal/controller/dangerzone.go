package controller

import (
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/service"
)

type DangerZoneController interface {
	GetAllActive() (*[]domain.DangerZone, error)
}

type dangerZoneController struct {
	svc service.DangerZoneService
}

func NewDangerZoneController(svc service.DangerZoneService) DangerZoneController {
	return &dangerZoneController{svc: svc}
}
func (ctrl *dangerZoneController) GetAllActive() (*[]domain.DangerZone, error) {
	return ctrl.svc.GetAllActive()
}
