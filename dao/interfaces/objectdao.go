package interfaces

import "github.com/alexyslozada/accounting-go/models"

type ObjectDAO interface {
	InsertObject(o *models.Object) error
	UpdateObject(o *models.Object) error
	DeleteObject(o *models.Object) error
	GetObjectByID(c string) (*models.Object, error)
	GetAllObject() ([]models.Object, error)
}
