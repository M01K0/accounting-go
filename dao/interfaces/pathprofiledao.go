package interfaces

import "github.com/alexyslozada/accounting-go/models"

type PathProfileDAO interface {
	Update(o *models.PathProfile) error
	GetByID(id int) (*models.PathProfile, error)
	GetByProfileID(id int16) ([]models.PathProfile, error)
}
