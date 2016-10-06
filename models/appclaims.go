package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// AppClaims provee una estructura personalizada para JWT claims
type AppClaims struct {
	User User `json:"user"`
	Scopes []Scope `json:"scopes"`
	jwt.StandardClaims
}
