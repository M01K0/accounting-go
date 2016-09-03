package models

import "time"

// CostCenter centro de costo
type CostCenter struct {
	ID         int16     `json:"id"`
	Code       string    `json:"code"`
	CostCenter string    `json:"costCenter"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
