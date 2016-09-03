package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountingHeaderDAO interface {
	Insert(o *models.AccountingHeader) error
	Update(o *models.AccountingHeader) error
	Delete(o *models.AccountingHeader) error
	GetByID(id int) (*models.AccountingHeader, error)
	GetAll() ([]models.AccountingHeader, error)
}
