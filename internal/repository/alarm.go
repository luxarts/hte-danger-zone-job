package repository

import "log"

type AlarmRepository interface {
	Send(deviceID string, message string) error
}

type alarmRepository struct {
}

func NewAlarmRepository() AlarmRepository {
	return &alarmRepository{}
}

func (repo *alarmRepository) Send(deviceID string, message string) error {
	log.Printf("%s->%s\n", deviceID, message)
	return nil
}
