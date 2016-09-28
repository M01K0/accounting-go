package models

import "time"

type PathProfile struct {
	ID        int   `json:"id"`
	ProfileID int16 `json:"profileId"`
	Path      `json:"path"`
	Post      bool      `json:"post"`
	Put       bool      `json:"put"`
	Del       bool      `json:"delete"`
	Get       bool      `json:"get"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
