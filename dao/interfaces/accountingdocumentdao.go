package interfaces

import "github.com/alexyslozada/accounting-go/models"

type AccountingDocumentDAO interface {
	Insert(o *models.AccountingDocument) error
	Update(o *models.AccountingDocument) error
	Delete(o *models.AccountingDocument) error
	GetByID(id int) (*models.AccountingDocument, error)
	GetAll() ([]models.AccountingDocument, error)
}
