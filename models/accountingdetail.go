package models

// AccountingDetail Detalle de movimiento contable
type AccountingDetail struct {
	ID                 int `json:"id"`
	AccountingHeaderID int `json:"idEncabezado"`
	AccountPUC         `json:"cuentaPuc"`
	Debit              float32 `json:"debito"`
	Credit             float32 `json:"credito"`
	ThirdParty         `json:"tercero"`
	CostCenter         `json:"centroCosto"`
}
