package models

import "time"

// AccountingDetail Detalle de movimiento contable
type AccountingDetail struct {
	ID                 int `json:"id"`
	AccountingHeaderID int `json:"accountingHeaderID"`
	AccountPUC         `json:"accountPUC"`
	Debit              float32 `json:"debit"`
	Credit             float32 `json:"credit"`
	ThirdParty         `json:"thirdParty"`
	CostCenter         `json:"costCenter"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
