package interfaces

import "github.com/alexyslozada/accounting-go/models"

// PathDAO Interface para el dao de Objeto
type PathDAO interface {
	Insert(o *models.Path) error
	Update(o *models.Path) error
	Delete(o *models.Path) error
	GetByID(id int) (*models.Path, error)
	GetAll() ([]models.Path, error)
}
