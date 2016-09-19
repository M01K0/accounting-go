package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountPeriodDAO interface {
	Insert(o *models.AccountPeriod) error
	Update(o *models.AccountPeriod) error
	Delete(o *models.AccountPeriod) error
	GetByID(id int) (*models.AccountPeriod, error)
	GetAll() ([]models.AccountPeriod, error)
}
