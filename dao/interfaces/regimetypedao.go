package interfaces

import "github.com/alexyslozada/accounting-go/models"

type RegimeTypeDAO interface {
	Insert(o *models.RegimeType) error
	Update(o *models.RegimeType) error
	Delete(o *models.RegimeType) error
	GetByID(id int) (*models.RegimeType, error)
	GetAll() ([]models.RegimeType, error)
}
