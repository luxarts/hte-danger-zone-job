package controller

import (
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"hte-danger-zone-job/internal/domain"
	"hte-danger-zone-job/internal/service"
	"log"
)

type GeolocController interface {
	Process(body *redis.XMessage) error
}

type geolocController struct {
	zoneSvc service.ZoneService
}

func NewGeolocController(zoneSvc service.ZoneService) GeolocController {
	return &geolocController{zoneSvc: zoneSvc}
}

func (ctrl *geolocController) Process(body *redis.XMessage) error {
	streamID := body.ID
	var userID string
	var p domain.Payload
	for k, v := range body.Values {
		userID = k
		err := json.Unmarshal([]byte(v.(string)), &p)
		if err != nil {
			return err
		}
	}

	log.Printf("-----------\n"+
		"Stream ID: %s\n"+
		"User ID: %s\n"+
		"Payload: %+v\n"+
		"-----------\n",
		streamID, userID, p)

	err := ctrl.zoneSvc.Verify(userID, &p)
	if err != nil {
		return err
	}

	return nil
}
