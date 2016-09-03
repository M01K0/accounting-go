package models

// TaxDetail Detalle del impuesto.
type TaxDetail struct {
	ID         int16   `json:"id"`
	TaxID      int16   `json:"idImpuesto"`
	Detail     string  `json:"detalle"`
	Percentage float32 `json:"porcentaje"`
	AccountID  int16   `json:"idCuenta"`
	Nature     string  `json:"naturaleza"`
	BaseValue  float32 `json:"valorBase"`
}
