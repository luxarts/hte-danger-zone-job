package service

import (
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
	"log"
	"math"
)

type ZoneService interface {
	Verify(userID string, p *domain.Payload) error
}

type zoneService struct {
	dzRepo    repository.DangerZoneRepository
	alarmRepo repository.AlarmRepository
}

func NewZoneService(dzRepo repository.DangerZoneRepository, alarmRepo repository.AlarmRepository) ZoneService {
	return &zoneService{
		dzRepo:    dzRepo,
		alarmRepo: alarmRepo,
	}
}

func (svc *zoneService) Verify(userID string, p *domain.Payload) error {
	dz, err := svc.dzRepo.GetByUserID(userID)
	if err != nil {
		return err
	}
	if dz == nil {
		return nil
	}

	log.Printf("zone-> lat:%f\tlon:%f\n", dz.CenterLatitude, dz.CenterLongitude)
	log.Printf("user-> lat:%f\tlon:%f\n", p.Latitude, p.Longitude)

	// Pythagorean Theorem: c = sqrt(a^2 + b^2)
	distance := math.Sqrt(math.Pow(p.Latitude-dz.CenterLatitude, 2) + math.Pow(p.Longitude-dz.CenterLongitude, 2))

	if distance >= dz.Radius && p.Timestamp <= dz.EndTimestamp { // Goes out before time
		err = svc.alarmRepo.Send(userID, defines.AlarmMessageOutsideZoneBeforeTime)
	} else if distance <= dz.Radius && p.Timestamp >= dz.EndTimestamp { // Inside after time
		err = svc.alarmRepo.Send(userID, defines.AlarmMessageInsideZoneAfterTime)
	}
	if err != nil {
		return err
	}

	return nil
}
