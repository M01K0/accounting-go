package models

// Object es la estructura de los objetos de seguridad
type Object struct {
	Code        string `json:"codigo"`
	Name        string `json:"nombre"`
	Description string `json:"descripcion"`
}
