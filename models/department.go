package models

import "time"

// Department departamento
type Department struct {
	ID         int16     `json:"id"`
	Code       string    `json:"code"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Cities     []City    `json:"cities"`
}
