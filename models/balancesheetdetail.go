package models

// BalanceSheetDetail Cierre contable - Detalle
type BalanceSheetDetail struct {
	ID                   int   `json:"id"`
	BalanceSheetHeaderID int16 `json:"idEncabezado"`
	AccountPUC           `json:"cuenta"`
	ThirdParty           `json:"tercero"`
	CostCenter           `json:"centroCosto"`
	PreviousBalance      float32 `json:"saldoAnterior"`
	Debit                float32 `json:"debito"`
	Credit               float32 `json:"credito"`
	ActualBalance        float32 `json:"saldoActual"`
}
