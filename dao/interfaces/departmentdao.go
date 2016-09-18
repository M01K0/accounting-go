package interfaces

import "github.com/alexyslozada/accounting-go/models"

type DepartmentDAO interface {
	Insert(o *models.Department) error
	Update(o *models.Department) error
	Delete(o *models.Department) error
	GetByID(id int) (*models.Department, error)
	GetAll() ([]models.Department, error)
	GetCities(o *models.Department) error
}
