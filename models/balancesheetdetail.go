package models

import "time"

// BalanceSheetDetail Cierre contable - Detalle
type BalanceSheetDetail struct {
	ID                   int   `json:"id"`
	BalanceSheetHeaderID int16 `json:"balanceSheetHeaderId"`
	AccountPUC           `json:"accountPUC"`
	ThirdParty           `json:"thirdParty"`
	CostCenter           `json:"costCenter"`
	PreviousBalance      float32   `json:"previousBalance"`
	Debit                float32   `json:"debit"`
	Credit               float32   `json:"credit"`
	CurrentBalance       float32   `json:"currentBalance"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
