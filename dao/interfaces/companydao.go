package interfaces

import "github.com/alexyslozada/accounting-go/models"

type CompanyDAO interface {
	Insert(o *models.Company) error
	Update(o *models.Company) error
	Delete(o *models.Company) error
	GetByID(id int) (*models.Company, error)
	GetAll() ([]models.Company, error)
}
