package models

import "time"

// FunctionaryType tipo de funcionario
type FunctionaryType struct {
	ID          int16     `json:"id"`
	Functionary string    `json:"functionary"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
