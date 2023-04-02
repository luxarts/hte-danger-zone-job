package repository

import "log"

type AlarmRepository interface {
	Send(userID string, message string) error
}

type alarmRepository struct {
}

func NewAlarmRepository() AlarmRepository {
	return &alarmRepository{}
}

func (repo *alarmRepository) Send(userID string, message string) error {
	log.Printf("%s->%s\n", userID, message)
	return nil
}
