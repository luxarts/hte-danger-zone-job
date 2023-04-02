package domain

type Payload struct {
	Timestamp int64   `json:"ts"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
