package models

import "time"

// AccountingDocument documento contable
type AccountingDocument struct {
	ID           int16     `json:"id"`
	Abbreviation string    `json:"abbreviation"`
	DocumentName string    `json:"documentName"`
	Consecutive  int       `json:"consecutive"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
