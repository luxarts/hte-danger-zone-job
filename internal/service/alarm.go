package service

import (
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
)

type AlarmService interface {
	Send(body *domain.SendAlarmReq) error
}

type alarmService struct {
	dzcRepo repository.DangerZoneCacheRepository
	repo    repository.AlarmRepository
}

func NewAlarmService(repo repository.AlarmRepository) AlarmService {
	return &alarmService{repo: repo}
}
func (svc *alarmService) Send(body *domain.SendAlarmReq) error {
	dzc, err := svc.dzcRepo.GetByDeviceID(body.DeviceID)
	if err != nil {
		return err
	}

	return svc.repo.Send(body.DeviceID, dzc.CompanyID, body.Message, dzc.Latitude, dzc.Longitude, dzc.CountryID)
}
