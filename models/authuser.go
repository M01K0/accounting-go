package models

type AuthUser struct {
	User        `json:"user"`
	TokenPost   string `json:"tokenPost"`
	TokenPut    string `json:"tokenPut"`
	TokenDelete string `json:"tokenDelete"`
	TokenGet    string `json:"tokenGet"`
}
