package interfaces

import "github.com/alexyslozada/accounting-go/models"

type PermissionsDAO interface {
	GetScopes(id int16) ([]models.Scope, error)
}
