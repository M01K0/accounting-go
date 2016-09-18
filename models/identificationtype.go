package models

import "time"

// IdentificationType tipo de identificación del tercero.
type IdentificationType struct {
	ID                 int16     `json:"id"`
	Initials           string    `json:"initials"`
	IdentificationName string    `json:"identificationName"`
	DianCode           string    `json:"dianCode"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
