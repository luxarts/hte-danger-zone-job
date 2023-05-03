package domain

type Alarm struct {
	AssetID   string        `json:"asset_id"`
	Type      string        `json:"type"`
	Action    string        `json:"action"`
	Timestamp int64         `json:"ts"`
	CompanyID string        `json:"company_id"`
	Position  AlarmPosition `json:"position"`
	Text      string        `json:"text"`
	CountryID int           `json:"country_id"`
	Device    AlarmDevice   `json:"device"`
}
type AlarmPosition struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type AlarmDevice struct {
	ID     string  `json:"_id"`
	Groups *string `json:"groups"`
}
