package models

// AccountingDocument documento contable
type AccountingDocument struct {
	ID           int16  `json:"id"`
	Abbreviation string `json:"abreviatura"`
	Document     string `json:"documento"`
	Consecutive  int    `json:"consecutivo"`
}
