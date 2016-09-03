package interfaces

import "github.com/alexyslozada/accounting-go/models"

type FunctionaryCompanyDAO interface {
	Insert(o *models.FunctionaryCompany) error
	Update(o *models.FunctionaryCompany) error
	Delete(o *models.FunctionaryCompany) error
	GetByID(id int) (*models.FunctionaryCompany, error)
	GetAll() ([]models.FunctionaryCompany, error)
}
