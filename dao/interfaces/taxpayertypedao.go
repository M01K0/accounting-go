package interfaces

import "github.com/alexyslozada/accounting-go/models"

type TaxpayerTypeDAO interface {
	Insert(o *models.TaxpayerType) error
	Update(o *models.TaxpayerType) error
	Delete(o *models.TaxpayerType) error
	GetByID(id int) (*models.TaxpayerType, error)
	GetAll() ([]models.TaxpayerType, error)
}
