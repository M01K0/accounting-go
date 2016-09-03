package interfaces

import "github.com/alexyslozada/accounting-go/models"

type TaxDAO interface {
	Insert(o *models.Tax) error
	Update(o *models.Tax) error
	Delete(o *models.Tax) error
	GetByID(id int) (*models.Tax, error)
	GetAll() ([]models.Tax, error)
}
