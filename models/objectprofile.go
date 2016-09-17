package models

import "time"

// ObjectByProfile son los objetos por perfil de la aplicación
type ObjectProfile struct {
	ID        int   `json:"id"`
	ProfileID int16 `json:"profileId"`
	Object    `json:"object"`
	Creates   bool      `json:"creates"`
	Modify    bool      `json:"modify"`
	Erase     bool      `json:"erase"`
	Query     bool      `json:"query"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
