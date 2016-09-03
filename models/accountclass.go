package models

import "time"

// AccountClass Clase de cuenta
type AccountClass struct {
	ID         int16  `json:"id"`
	ReportType `json:"reportType"`
	AccountClass      string `json:"accountClass"`
	Nature     string `json:"nature"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
