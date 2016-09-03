package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountingDetailDAO interface {
	Insert(o *models.AccountingDetail) error
	DeleteByHeader(i int) error
	GetByID(id int) (*models.AccountingDetail, error)
	GetAll() ([]models.AccountingDetail, error)
}
