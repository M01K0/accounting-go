package interfaces

import "github.com/alexyslozada/accounting-go/models"

// ProfileDAO Interface para el dao del perfil
type ProfileDAO interface {
	Insert(p *models.Profile) error
	Update(p *models.Profile) error
	Delete(p *models.Profile) error
	GetByID(id int16) (*models.Profile, error)
	GetAll() ([]models.Profile, error)
}
