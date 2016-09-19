package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountLevelDAO interface {
	Insert(o *models.AccountLevel) error
	Update(o *models.AccountLevel) error
	Delete(o *models.AccountLevel) error
	GetByID(id int) (*models.AccountLevel, error)
	GetAll() ([]models.AccountLevel, error)
}
