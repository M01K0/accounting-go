package interfaces

import "github.com/alexyslozada/accounting-go/models"

type TaxDetailDAO interface {
	Insert(o *models.TaxDetail) error
	Update(o *models.TaxDetail) error
	Delete(o *models.TaxDetail) error
	GetByID(id int) (*models.TaxDetail, error)
	GetAll() ([]models.TaxDetail, error)
}
