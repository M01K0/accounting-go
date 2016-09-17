package interfaces

import "github.com/alexyslozada/accounting-go/models"

type ObjectProfileDAO interface {
	Update(o *models.ObjectProfile) error
	GetByID(id int) (*models.ObjectProfile, error)
	GetByProfileID(id int16) ([]models.ObjectProfile, error)
}
