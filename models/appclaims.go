package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// AppClaims provee una estructura personalizada para JWT claims
type AppClaims struct {
	User User `json:"user"`
	Scopes map[string][]string `json:"scopes"`
	jwt.StandardClaims
}
