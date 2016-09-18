package interfaces

import "github.com/alexyslozada/accounting-go/models"

type UserDAO interface {
	Insert(o *models.User) error
	Update(o *models.User) error
	Delete(o *models.User) error
	GetByID(id int) (*models.User, error)
	GetAll() ([]models.User, error)
}
