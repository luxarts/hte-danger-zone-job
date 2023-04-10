package domain

type DangerZone struct {
	DeviceID     string  `json:"did"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lon"`
	Radius       float64 `json:"r"`
	EndTimestamp int64   `json:"e_ts"`
}
