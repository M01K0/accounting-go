package models

import "time"

// User usuario del sistema
type User struct {
	ID             int16  `json:"id"`
	Identification string `json:"identification"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Passwd         string `json:"passwd"`
	Profile        `json:"profile"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
