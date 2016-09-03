package models

import "time"

// AccountingHeader Encabezado de movimiento contable
type AccountingHeader struct {
	ID                 int `json:"id"`
	AccountingDocument `json:"accountingDocument"`
	Consecutive        int       `json:"consecutive"`
	MovementDate       time.Time `json:"movementDate"`
	Commentary         string    `json:"commentary"`
	AccountPeriod      `json:"accountPeriod"`
	UserCreater        User               `json:"userCreater"`
	UserUpdater        User               `json:"userUpdater"`
	Anulled            bool               `json:"anulled"`
	CreatedAt          time.Time          `json:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt"`
	Details            []AccountingDetail `json:"details"`
}
