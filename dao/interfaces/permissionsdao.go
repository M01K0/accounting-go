package interfaces

import "github.com/alexyslozada/accounting-go/models"

type PermissionsDAO interface {
	IsPermitted(profile models.Profile, path string, method string) (bool, error)
}
