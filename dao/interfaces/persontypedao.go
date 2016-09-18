package interfaces

import "github.com/alexyslozada/accounting-go/models"

type PersonTypeDAO interface {
	Insert(o *models.PersonType) error
	Update(o *models.PersonType) error
	Delete(o *models.PersonType) error
	GetByID(id int) (*models.PersonType, error)
	GetAll() ([]models.PersonType, error)
}
