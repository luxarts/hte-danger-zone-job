package domain

type DangerZone struct {
	DeviceID     string  `json:"device_id"`
	CompanyID    string  `json:"company_id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Radius       float64 `json:"radius"`
	EndTimestamp int64   `json:"end_ts"`
	CountryID    int     `json:"country_id"`
}
