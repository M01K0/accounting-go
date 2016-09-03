package interfaces

import "github.com/alexyslozada/accounting-go/models"

// ObjectDAO Interface para el dao de Objeto
type ObjectDAO interface {
	Insert(o *models.Object) error
	Update(o *models.Object) error
	Delete(o *models.Object) error
	GetByID(id int) (*models.Object, error)
	GetAll() ([]models.Object, error)
}
