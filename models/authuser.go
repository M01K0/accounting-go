package models

type AuthUser struct {
	User `json:"user"`
	Token string `json:"token"`
}
