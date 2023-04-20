package domain

type DangerZone struct {
	DeviceID     string  `json:"device_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       float64 `json:"radius"`
	EndTimestamp int64   `json:"end_timestamp"`
}
