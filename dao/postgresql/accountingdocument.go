package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type AccountingDocumentDAOPsql struct {}

// Insert insertar registro en la BD
func (dao AccountingDocumentDAOPsql) Insert(obj *models.AccountingDocument) error {
	query := "INSERT INTO accounting_document (abbreviation, document_name) VALUES (upper($1), upper($2)) RETURNING id, abbreviation, document_name, consecutive, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Abbreviation, obj.DocumentName)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao AccountingDocumentDAOPsql) Update(obj *models.AccountingDocument) error {
	query := "UPDATE accounting_document SET abbreviation = upper($2), document_name = upper($3), updated_at = now() WHERE id = $1 RETURNING id, abbreviation, document_name, consecutive, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Abbreviation, obj.DocumentName)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao AccountingDocumentDAOPsql) Delete(obj *models.AccountingDocument) error {
	query := "DELETE FROM accounting_document WHERE id = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	result, err := stmt.Exec(obj.ID)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
			return errors.New("No se eliminó ningún registro")
	}
	obj = new(models.AccountingDocument)
	return nil
}

// GetByID consultar registro por id
func (dao AccountingDocumentDAOPsql) GetByID(id int) (*models.AccountingDocument, error) {
	query := "SELECT id, abbreviation, document_name, consecutive, created_at, updated_at FROM accounting_document WHERE id = $1"
	obj := &models.AccountingDocument{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = dao.rowToObject(row, obj)
	return obj, err
}

// GetAll Consulta todos los registros de la bd
func (dao AccountingDocumentDAOPsql) GetAll() ([]models.AccountingDocument, error) {
	query := "SELECT id, abbreviation, document_name, consecutive, created_at, updated_at FROM accounting_document ORDER BY id"
	objs := make([]models.AccountingDocument, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.AccountingDocument
		err = rows.Scan(&o.ID, &o.Abbreviation, &o.DocumentName, &o.Consecutive, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountingDocumentDAOPsql) rowToObject(row *sql.Row, o *models.AccountingDocument) error {
	return row.Scan(&o.ID, &o.Abbreviation, &o.DocumentName, &o.Consecutive, &o.CreatedAt, &o.UpdatedAt)
}
