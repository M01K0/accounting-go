package controllers

import "github.com/alexyslozada/accounting-go/models"

type (
	AuthUserResource struct {
		Data models.AuthUser `json:"data"`
	}
	CostCenterResource struct {
		Data models.CostCenter `json:"data"`
	}
	CostCentersResource struct {
		Data []models.CostCenter `json:"data"`
	}
	LoginResource struct {
		Data models.Login `json:"data"`
	}
)
