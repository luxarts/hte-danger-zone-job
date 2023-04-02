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
	userID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dz := domain.DangerZone{
		CenterLatitude:  2,
		CenterLongitude: 0.5,
		Radius:          1.5,
		EndTimestamp:    10,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(&dz, nil)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Nil(t, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_OutsideZoneAfterTime_NoAlarm_OK(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 11,
		Latitude:  3.2,
		Longitude: 1.5,
	}

	dz := domain.DangerZone{
		CenterLatitude:  2,
		CenterLongitude: 0.5,
		Radius:          1.5,
		EndTimestamp:    10,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(&dz, nil)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Nil(t, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_InsideZoneAfterTime_Alarm_OK(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 11,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dz := domain.DangerZone{
		CenterLatitude:  2,
		CenterLongitude: 0.5,
		Radius:          1.5,
		EndTimestamp:    10,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(&dz, nil)

	alarmRepo := new(repository.MockAlarmRepository)
	alarmRepo.On("Send", userID, defines.AlarmMessageInsideZoneAfterTime).Return(nil)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Nil(t, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_OutsideZoneBeforeTime_Alarm_OK(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.4,
	}

	dz := domain.DangerZone{
		CenterLatitude:  2,
		CenterLongitude: 0.5,
		Radius:          1.5,
		EndTimestamp:    10,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(&dz, nil)

	alarmRepo := new(repository.MockAlarmRepository)
	alarmRepo.On("Send", userID, defines.AlarmMessageOutsideZoneBeforeTime).Return(nil)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Nil(t, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_InsideZoneBeforeTime_NoAlarm_ErrorGettingDangerZone(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dzErr := errors.New("error getting danger zone")
	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(nil, dzErr)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Equal(t, dzErr, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_InsideZoneBeforeTime_NoAlarm_NoDangerZone(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(nil, nil)

	alarmRepo := new(repository.MockAlarmRepository)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Nil(t, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_InsideZoneAfterTime_Alarm_Error(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 11,
		Latitude:  3.2,
		Longitude: 1.3,
	}

	dz := domain.DangerZone{
		CenterLatitude:  2,
		CenterLongitude: 0.5,
		Radius:          1.5,
		EndTimestamp:    10,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(&dz, nil)

	alarmErr := errors.New("error sending alarm")
	alarmRepo := new(repository.MockAlarmRepository)
	alarmRepo.On("Send", userID, defines.AlarmMessageInsideZoneAfterTime).Return(alarmErr)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Equal(t, alarmErr, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
func TestZoneService_Verify_OutsideZoneBeforeTime_Alarm_Error(t *testing.T) {
	// Given
	userID := "123"
	p := &domain.Payload{
		Timestamp: 1,
		Latitude:  3.2,
		Longitude: 1.4,
	}

	dz := domain.DangerZone{
		CenterLatitude:  2,
		CenterLongitude: 0.5,
		Radius:          1.5,
		EndTimestamp:    10,
	}

	dzRepo := new(repository.MockDangerZoneRepository)
	dzRepo.On("GetByUserID", userID).Return(&dz, nil)

	alarmErr := errors.New("error sending alarm")
	alarmRepo := new(repository.MockAlarmRepository)
	alarmRepo.On("Send", userID, defines.AlarmMessageOutsideZoneBeforeTime).Return(alarmErr)

	svc := NewZoneService(dzRepo, alarmRepo)

	// When
	err := svc.Verify(userID, p)

	// Then
	assert.Equal(t, alarmErr, err)
	dzRepo.AssertExpectations(t)
	alarmRepo.AssertExpectations(t)
}
