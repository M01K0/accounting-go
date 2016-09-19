package models

import "time"

// FunctionaryCompany Funcionario de la empresa.
type FunctionaryCompany struct {
	ID                   int16 `json:"id"`
	FunctionaryType      `json:"functionaryType"`
	IdentificationType   `json:"identificationType"`
	IdentificationNumber string    `json:"identificationNumber"`
	VerificationDigit    string    `json:"verificationDigit"`
	Functionary          string    `json:"functionary"`
	Active               bool      `json:"active"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
