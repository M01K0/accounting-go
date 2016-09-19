package models

import "time"

// ReportType Tipo de informe
type ReportType struct {
	ID        int16     `json:"id"`
	Report    string    `json:"informe"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
