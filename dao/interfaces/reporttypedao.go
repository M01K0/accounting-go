package interfaces

import "github.com/alexyslozada/accounting-go/models"

type ReportTypeDAO interface {
	Insert(o *models.ReportType) error
	Update(o *models.ReportType) error
	Delete(o *models.ReportType) error
	GetByID(id int) (*models.ReportType, error)
	GetAll() ([]models.ReportType, error)
}
