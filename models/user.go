package models

// User usuario del sistema
type User struct {
	ID             int16  `json:"id"`
	Identification string `json:"identificacion"`
	Name           string `json:"nombre"`
	Email          string `json:"correo"`
	Password       string `json:"clave"`
	Profile        `json:"perfil"`
	Active         bool `json:"activo"`
}
