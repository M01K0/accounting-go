package models

import "time"

// TaxpayerType Tipo de contribuyente
type TaxpayerType struct {
	ID   int16  `json:"id"`
	Taxpayer string `json:"taxpayer"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
