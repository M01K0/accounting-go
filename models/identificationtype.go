package models

// IdentificationType tipo de identificación del tercero.
type IdentificationType struct {
	ID       int16  `json:"id"`
	Initials string `json:"sigla"`
	Document string `json:"documento"`
	DianCode string `json:"codigoDian"`
}
