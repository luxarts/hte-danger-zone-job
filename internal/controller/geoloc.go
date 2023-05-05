package controller

import (
	"encoding/json"
	"hte-danger-zone-job/internal/defines"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/service"
	"log"
)

type GeolocController interface {
	Process(deviceID string, body string) (*defines.ResponseStatus, error)
}

type geolocController struct {
	zoneSvc service.ZoneService
}

func NewGeolocController(zoneSvc service.ZoneService) GeolocController {
	return &geolocController{zoneSvc: zoneSvc}
}

func (ctrl *geolocController) Process(deviceID string, body string) (*defines.ResponseStatus, error) {
	var p domain.Payload

	err := json.Unmarshal([]byte(body), &p)
	if err != nil {
		return nil, err
	}

	log.Printf("Device ID: %s\tPayload: %+v\n", deviceID, p)

	return ctrl.zoneSvc.Verify(deviceID, &p)
}
