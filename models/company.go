package models

import "time"

// Company Datos generales de la empresa
type Company struct {
	ID                   int16 `json:"id"`
	IdentificationType   `json:"identificationType"`
	IdentificationNumber string `json:"identificationNumber"`
	VerificationDigit    string `json:"verificationDigit"`
	Company              string `json:"company"`
	Address              string `json:"address"`
	Phone                string `json:"phone"`
	Department           `json:"department"`
	City                 `json:"city"`
	Web                  string `json:"web"`
	Email                string `json:"email"`
	Activity             string `json:"activity"`
	AutoRretenedor       bool   `json:"autorretenedor"`
	PersonType           `json:"personType"`
	RegimeType           `json:"regimeType"`
	TaxpayerType         `json:"taxpayerType"`
	Logo                 string    `json:"logo"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt`
}
