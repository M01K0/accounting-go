package models

// CostCenter centro de costo
type CostCenter struct {
	ID   int16  `json:"id"`
	Code string `json:"codigo"`
	Name string `json:"nombre"`
}
