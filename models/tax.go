package models

import "time"

// Tax Impuesto. Es el control de los impuestos.
type Tax struct {
	ID   int16  `json:"id"`
	Tax  string `json:"nombre"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
