package models

import "time"

// TaxDetail Detalle del impuesto.
type TaxDetail struct {
	ID         int16 `json:"id"`
	Tax        `json:"tax"`
	Detail     string  `json:"detail"`
	Percentage float32 `json:"porcentaje"`
	AccountPUC `json:"accountPUC"`
	Nature     string    `json:"nature"`
	BaseValue  float32   `json:"baseValue"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
