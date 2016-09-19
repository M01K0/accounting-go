package models

import "time"

// AccountLevel nivel de cuenta
type AccountLevel struct {
	ID           int16     `json:"id"`
	AccountLevel string    `json:"accountLevel"`
	Digits       int8      `json:"digits"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
