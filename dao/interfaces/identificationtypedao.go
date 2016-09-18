package interfaces

import "github.com/alexyslozada/accounting-go/models"

type IdentificationTypeDAO interface {
	Insert(o *models.IdentificationType) error
	Update(o *models.IdentificationType) error
	Delete(o *models.IdentificationType) error
	GetByID(id int) (*models.IdentificationType, error)
	GetAll() ([]models.IdentificationType, error)
}
