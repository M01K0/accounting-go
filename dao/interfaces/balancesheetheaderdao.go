package interfaces

import "github.com/alexyslozada/accounting-go/models"

type BalanceSheetHeaderDAO interface {
	Insert(o *models.BalanceSheetHeader) error
	Update(o *models.BalanceSheetHeader) error
	Delete(o *models.BalanceSheetHeader) error
	GetByID(id int) (*models.BalanceSheetHeader, error)
	GetAll() ([]models.BalanceSheetHeader, error)
}
