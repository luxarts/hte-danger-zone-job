package controller

import (
	"encoding/json"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/service"
	"log"
)

type GeolocController interface {
	Process(deviceID string, body string) error
}

type geolocController struct {
	zoneSvc service.ZoneService
}

func NewGeolocController(zoneSvc service.ZoneService) GeolocController {
	return &geolocController{zoneSvc: zoneSvc}
}

func (ctrl *geolocController) Process(deviceID string, body string) error {
	var p domain.Payload

	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return err
	}

	log.Printf("Device ID: %s\tPayload: %+v\n", deviceID, p)

	err = ctrl.zoneSvc.Verify(deviceID, &p)
	if err != nil {
		return err
	}

	return nil
}
