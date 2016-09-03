package interfaces

import "github.com/alexyslozada/accounting-go/models"

type CityDAO interface {
	Insert(o *models.City) error
	Update(o *models.City) error
	Delete(o *models.City) error
	GetByID(id int) (*models.City, error)
	GetAll() ([]models.City, error)
}
