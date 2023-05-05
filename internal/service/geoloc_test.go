package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/repository"
	"testing"
)

func TestZoneService_Verify_InsideZoneBeforeTime_NoAlarm_OK(t *testing.T) {
	// Given
	deviceID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dz := domain.DangerZone{
		Latitude:     2,
		Longitude:    0.5,
		Radius:       1.5,
		EndTimestamp: 10,
	}

	dzcRepo := new(repository.MockDangerZoneCacheRepository)
	dzcRepo.On("GetByDeviceID", deviceID).Return(&dz, nil)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzcRepo, alarmRepo)

	// When
	resp, err := svc.Verify(deviceID, p)

	// Then
	assert.Nil(t, err)
	assert.Nil(t, resp)
	dzcRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_OutsideZoneBeforeTime_Alarm_OK(t *testing.T) {
	// Given
	deviceID := "someDevice"

	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.4,
	}

	dz := domain.DangerZone{
		CompanyID:    "someCompany",
		Latitude:     2,
		Longitude:    0.5,
		Radius:       1.5,
		EndTimestamp: 10,
		CountryID:    1,
	}

	dzcRepo := new(repository.MockDangerZoneCacheRepository)
	dzcRepo.On("GetByDeviceID", deviceID).Return(&dz, nil)

	alarmRepo := new(repository.MockAlarmRepository)
	alarmRepo.On("Send", deviceID, dz.CompanyID, defines.AlarmMessageOutsideZoneBeforeTime, p.Latitude, p.Longitude, dz.CountryID).Return(nil)

	svc := NewZoneService(dzcRepo, alarmRepo)

	// When
	resp, err := svc.Verify(deviceID, p)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, defines.ResponseStatusOutBeforeTime, *resp)
	dzcRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_ErrorGettingDangerZone(t *testing.T) {
	// Given
	deviceID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dzErr := errors.New("error getting danger zone")
	dzcRepo := new(repository.MockDangerZoneCacheRepository)
	dzcRepo.On("GetByDeviceID", deviceID).Return(nil, dzErr)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzcRepo, alarmRepo)

	// When
	resp, err := svc.Verify(deviceID, p)

	// Then
	assert.Equal(t, dzErr, err)
	assert.Nil(t, resp)
	dzcRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_NoDangerZone(t *testing.T) {
	// Given
	deviceID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dzcRepo := new(repository.MockDangerZoneCacheRepository)
	dzcRepo.On("GetByDeviceID", deviceID).Return(nil, nil)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzcRepo, alarmRepo)

	// When
	resp, err := svc.Verify(deviceID, p)

	// Then
	assert.Nil(t, err)
	assert.Nil(t, resp)
	dzcRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_OutsideZoneBeforeTime_Alarm_Error(t *testing.T) {
	// Given
	deviceID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.4,
	}

	dz := domain.DangerZone{
		CompanyID:    "someCompany",
		Latitude:     2,
		Longitude:    0.5,
		Radius:       1.5,
		EndTimestamp: 10,
		CountryID:    1,
	}

	dzcRepo := new(repository.MockDangerZoneCacheRepository)
	dzcRepo.On("GetByDeviceID", deviceID).Return(&dz, nil)

	alarmErr := errors.New("error sending alarm")
	alarmRepo := new(repository.MockAlarmRepository)
	alarmRepo.On("Send", deviceID, dz.CompanyID, defines.AlarmMessageOutsideZoneBeforeTime, p.Latitude, p.Longitude, dz.CountryID).Return(alarmErr)

	svc := NewZoneService(dzcRepo, alarmRepo)

	// When
	resp, err := svc.Verify(deviceID, p)

	// Then
	assert.Equal(t, alarmErr, err)
	assert.Nil(t, resp)
	dzcRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
