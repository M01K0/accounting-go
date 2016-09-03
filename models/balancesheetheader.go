package models

import "time"

// BalanceSheetHeader Cierre contable - Encabezado
type BalanceSheetHeader struct {
	ID            int16 `json:"id"`
	User          `json:"user"`
	AccountPeriod `json:"accountPeriod"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
