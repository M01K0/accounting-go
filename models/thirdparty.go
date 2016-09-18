package models

import "time"

// ThirdParty Terceros
type ThirdParty struct {
	ID                   int16 `json:"id"`
	IdentificationType   `json:"identificationType"`
	IdentificationNumber string `json:"identificationNumber"`
	VerificationDigit    string `json:"verificationDigit"`
	PersonType           `json:"personType"`
	RegimeType           `json:"regimeType"`
	TaxpayerType         `json:"taxpayerType"`
	BusinessName         string `json:"businessName"`
	LastName             string `json:"lastName"`
	SecondLastName       string `json:"secondLastName"`
	FirstName            string `json:"firstName"`
	MiddleName           string `json:"middleName"`
	Address              string `json:"address"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Department           `json:"department"`
	City                 `json:"city"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
