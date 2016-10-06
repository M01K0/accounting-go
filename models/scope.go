package models

// Scope provee una estructura para los permisos del perfil
type Scope struct {
	Path string `json:"path"`
	Methods []string `json:"methods"`
}
