package interfaces

import "github.com/alexyslozada/accounting-go/models"

type CostCenterDAO interface {
	Insert(o *models.CostCenter) error
	Update(o *models.CostCenter) error
	Delete(id int16) error
	GetByID(id int16) (*models.CostCenter, error)
	GetAll() ([]models.CostCenter, error)
	GetAllPagination(p int, l int, o int, t string) ([]models.CostCenter, error)
}
