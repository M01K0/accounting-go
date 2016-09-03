package models

import "time"

// BalanceSheetHeader Cierre contable - Encabezado
type BalanceSheetHeader struct {
	ID            int16     `json:"id"`
	SystemDate    time.Time `json:"fechaSistema"`
	User          `json:"usuario"`
	AccountPeriod `json:"periodo"`
}
