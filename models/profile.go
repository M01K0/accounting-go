package models

// Profile es el perfil de la aplicación
type Profile struct {
	ID     int16  `json:"id"`
	Name   string `json:"nombre"`
	Active bool   `json:"activo"`
}
