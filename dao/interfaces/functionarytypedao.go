package interfaces

import "github.com/alexyslozada/accounting-go/models"

type FunctionaryTypeDAO interface {
	Insert(o *models.FunctionaryType) error
	Update(o *models.FunctionaryType) error
	Delete(o *models.FunctionaryType) error
	GetByID(id int) (*models.FunctionaryType, error)
	GetAll() ([]models.FunctionaryType, error)
}
