package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountClassDAO interface {
	Insert(o *models.AccountClass) error
	Update(o *models.AccountClass) error
	Delete(o *models.AccountClass) error
	GetByID(id int) (*models.AccountClass, error)
	GetAll() ([]models.AccountClass, error)
}
