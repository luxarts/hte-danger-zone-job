package domain

type DangerZone struct {
	CenterLatitude  float64 `json:"lat"`
	CenterLongitude float64 `json:"lon"`
	Radius          float64 `json:"r"`
	EndTimestamp    int64   `json:"end_ts"`
}
