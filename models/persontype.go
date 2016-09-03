package models

import "time"

// PersonType Tipo de persona ante la ley
type PersonType struct {
	ID        int16     `json:"id"`
	Person    string    `json:"person"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
