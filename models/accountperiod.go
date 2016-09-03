package models

import "time"

// AccountPeriod periodo contable
type AccountPeriod struct {
	ID        int16     `json:"id"`
	Year      int16     `json:"year"`
	Month     int8      `json:"month"`
	Open      bool      `json:"isOpen"`
	CloseDate time.Time `json:"closeDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
