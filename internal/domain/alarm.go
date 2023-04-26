package domain

type Alarm struct {
	AssetID   string  `json:"asset_id"`
	Type      string  `json:"type"`
	Action    string  `json:"action"`
	Timestamp int64   `json:"ts"`
	CompanyID string  `json:"company_id"`
	Position  *string `json:"position"`
	Text      string  `json:"text"`
	CountryID int     `json:"country_id"`
	Device    string  `json:"device"`
}
