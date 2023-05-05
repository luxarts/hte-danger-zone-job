package controller

import (
	"errors"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/service"
)

type AlarmController interface {
	Send(body *domain.SendAlarmReq) error
}

type alarmController struct {
	svc service.AlarmService
}

func NewAlarmController(svc service.AlarmService) AlarmController {
	return &alarmController{svc: svc}
}
func (ctrl *alarmController) Send(body *domain.SendAlarmReq) error {
	if !body.IsValid() {
		return errors.New("invalid body")
	}

	return ctrl.svc.Send(body)
}
