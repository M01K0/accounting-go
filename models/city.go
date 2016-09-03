package models

import "time"

// City ciudades de colombia
type City struct {
	ID           int16     `json:"id"`
	DepartmentID int16     `json:"departmentId"`
	Code         string    `json:"code"`
	City         string    `json:"city"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
