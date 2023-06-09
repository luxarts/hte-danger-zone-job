package service

import (
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
	"math"
)

type ZoneService interface {
	Verify(deviceID string, p *domain.Payload) (bool, error)
}

type zoneService struct {
	dzcRepo   repository.DangerZoneCacheRepository
	alarmRepo repository.AlarmRepository
}

func NewZoneService(dzcRepo repository.DangerZoneCacheRepository, alarmRepo repository.AlarmRepository) ZoneService {
	return &zoneService{
		dzcRepo:   dzcRepo,
		alarmRepo: alarmRepo,
	}
}

func (svc *zoneService) Verify(deviceID string, p *domain.Payload) (bool, error) {
	dz, err := svc.dzcRepo.GetByDeviceID(deviceID)
	if err != nil {
		return false, err
	}
	if dz == nil {
		return true, nil
	}

	distance := math.Sqrt(math.Pow(p.Latitude-dz.Latitude, 2) + math.Pow(p.Longitude-dz.Longitude, 2))
	var exit bool
	if distance >= dz.Radius && p.Timestamp <= dz.EndTimestamp { // Goes out before time
		err = svc.alarmRepo.Send(deviceID, "", defines.AlarmMessageOutsideZoneBeforeTime)
		exit = true
	} else if distance <= dz.Radius && p.Timestamp >= dz.EndTimestamp { // Inside after time
		err = svc.alarmRepo.Send(deviceID, "", defines.AlarmMessageInsideZoneAfterTime)
		exit = true
	}

	if err != nil {
		return false, err
	}

	return exit, nil
}
