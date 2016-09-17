package models

import "time"

// Object es la estructura de los objetos de seguridad
type Object struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	ObjectName  string    `json:"objectName"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
