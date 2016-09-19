package models

import "time"

// AccountPUC cuenta del plan único de cuentas
type AccountPUC struct {
	ID              int16  `json:"id"`
	Account         string `json:"account"`
	AccountName     string `json:"accountName"`
	AccountParentID int16  `json:"accountParent"`
	AccountClass    `json:"accountClass"`
	AccountLevel    `json:"accountLevel"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
