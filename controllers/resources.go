package controllers

import "github.com/alexyslozada/accounting-go/models"

type (
	AuthUserResource struct {
		Data models.AuthUser `json:"data"`
	}
	LoginResource struct {
		Data models.Login `json:"data"`
	}
)
