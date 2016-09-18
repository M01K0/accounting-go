package interfaces

import "github.com/alexyslozada/accounting-go/models"

type ThirdPartyDAO interface {
	Insert(o *models.ThirdParty) error
	Update(o *models.ThirdParty) error
	Delete(o *models.ThirdParty) error
	GetByID(id int) (*models.ThirdParty, error)
	GetAll() ([]models.ThirdParty, error)
}
