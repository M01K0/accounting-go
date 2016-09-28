package models

import "time"

type Path struct {
	ID          int       `json:"id"`
	Path        string    `json:"path"`
	PathName    string    `json:"pathName"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
