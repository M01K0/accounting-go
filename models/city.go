package models

// City ciudades de colombia
type City struct {
	ID              int16  `json:"id"`
	CodeDepartament string `json:"codigoDepartamento"`
	Code            string `json:"codigo"`
	Name            string `json:"nombre"`
}
