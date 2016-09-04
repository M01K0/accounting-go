package interfaces

import "github.com/alexyslozada/accounting-go/models"

// ProfileDAO Interface para el dao del perfil
type ProfileDAO interface {
	InsertProfile(p *models.Profile) error
	UpdateProfile(p *models.Profile) error
	DeleteProfile(p *models.Profile) error
	GetProfileByID(id int16) (*models.Profile, error)
	GetAllProfiles() ([]models.Profile, error)
}
