package models

import "time"

// AccountingHeader Encabezado de movimiento contable
type AccountingHeader struct {
	ID                 int `json:"id"`
	AccountingDocument `json:"documento"`
	Consecutive        int       `json:"consecutivo"`
	MovementDate       time.Time `json:"fechaMovimiento"`
	Comment            string    `json:"comentario"`
	AccountPeriod      `json:"periodo"`
	UserCreate         User      `json:"usuario"`
	CreatedAt          time.Time `json:"fechaCreacion"`
	UserUpdate         User      `json:"usuarioActualizacion"`
	UpdatedAt          time.Time `json:"fechaActualizacion"`
	Anulled            bool      `json:"anulado"`
}
