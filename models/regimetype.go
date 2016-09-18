package models

import "time"

// RegimeType tipo de régimen
type RegimeType struct {
	ID        int16     `json:"id"`
	Regime    string    `json:"regime"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
