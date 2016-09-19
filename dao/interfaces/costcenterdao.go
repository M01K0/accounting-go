package interfaces

import "github.com/alexyslozada/accounting-go/models"

type CostCenterDAO interface {
	Insert(o *models.CostCenter) error
	Update(o *models.CostCenter) error
	Delete(o *models.CostCenter) error
	GetByID(id int) (*models.CostCenter, error)
	GetAll() ([]models.CostCenter, error)
}
