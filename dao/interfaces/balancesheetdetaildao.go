package interfaces

import "github.com/alexyslozada/accounting-go/models"

type BalanceSheetDetailDAO interface {
	Insert(o *models.BalanceSheetDetail) error
	GetByID(id int) (*models.BalanceSheetDetail, error)
	GetAll() ([]models.BalanceSheetDetail, error)
}
