package models

// ObjectByProfile son los objetos por perfil de la aplicación
type ObjectByProfile struct {
	ID      int16 `json:"id"`
	Profile `json:"perfil"`
	Object  `json:"objeto"`
	Insert  bool `json:"insertar"`
	Update  bool `json:"modificar"`
	Delete  bool `json:"borrar"`
	Query   bool `json:"consultar"`
}
