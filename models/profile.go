package models

import "time"

// Profile es el perfil de la aplicación
type Profile struct {
	ID             int16           `json:"id"`
	Profile        string          `json:"profile"`
	Active         bool            `json:"active"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	ObjectsProfile []ObjectProfile `json:"objectsProfile"`
}
