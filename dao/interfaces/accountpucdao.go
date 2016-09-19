package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountPUCDAO interface {
	Insert(o *models.AccountPUC) error
	Update(o *models.AccountPUC) error
	Delete(o *models.AccountPUC) error
	GetByID(id int) (*models.AccountPUC, error)
	GetByAccount(a string) (*models.AccountPUC, error)
	GetAll() ([]models.AccountPUC, error)
}
