package interfaces

import "github.com/alexyslozada/accounting-go/models"

type LoginDAO interface {
	Login(u *models.User) error
}
