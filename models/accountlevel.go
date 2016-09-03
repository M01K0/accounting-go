package models

// AccountLevel nivel de cuenta
type AccountLevel struct {
	ID     int16  `json:"id"`
	Name   string `json:"nombre"`
	Digits int8   `json:"digitos"`
}
