package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type AccountingHeaderDAOPsql struct {}

// Insert insertar registro en la BD
func (dao AccountingHeaderDAOPsql) Insert(obj *models.AccountingHeader) error {
	query := `INSERT INTO accounting_headers (accounting_document_id, movement_date, commentary, account_period_id, user_id)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id, accounting_document_id, consecutive, movement_date, commentary, account_period_id, user_id, updated_user_id, anulled, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.AccountingDocument.ID, obj.MovementDate, obj.Commentary, obj.AccountPeriod.ID, obj.UserCreater.ID)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao AccountingHeaderDAOPsql) Update(obj *models.AccountingHeader) error {
	query := `UPDATE accounting_headers SET movement_date = $2, commentary = $3, account_period_id = $4, updated_user_id = $5, updated_at = now()
				WHERE id = $1
				RETURNING id, accounting_document_id, consecutive, movement_date, commentary, account_period_id, user_id, updated_user_id, anulled, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.MovementDate, obj.Commentary, obj.AccountPeriod.ID, obj.UserUpdater.ID)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao AccountingHeaderDAOPsql) Delete(obj *models.AccountingHeader) error {
	query := "DELETE FROM accounting_headers WHERE id = $1"
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
	obj = new(models.AccountingHeader)
	return nil
}

// GetByID consultar registro por id
func (dao AccountingHeaderDAOPsql) GetByID(id int) (*models.AccountingHeader, error) {
	query := `SELECT id, accounting_document_id, consecutive, movement_date, commentary, account_period_id, user_id, updated_user_id, anulled, created_at, updated_at
				FROM accounting_headers
				WHERE id = $1`
	obj := &models.AccountingHeader{}
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
func (dao AccountingHeaderDAOPsql) GetAll() ([]models.AccountingHeader, error) {
	query := `SELECT id, accounting_document_id, consecutive, movement_date, commentary, account_period_id, user_id, updated_user_id, anulled, created_at, updated_at
				FROM accounting_headers
				ORDER BY id`
	objs := make([]models.AccountingHeader, 0)
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
		var o models.AccountingHeader
		err = rows.Scan(&o.ID, &o.AccountingDocument.ID, &o.Consecutive, &o.MovementDate, &o.Commentary, &o.AccountPeriod.ID, &o.UserCreater.ID, &o.UserUpdater.ID, &o.Anulled, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountingHeaderDAOPsql) rowToObject(row *sql.Row, o *models.AccountingHeader) error {
	return row.Scan(&o.ID, &o.AccountingDocument.ID, &o.Consecutive, &o.MovementDate, &o.Commentary, &o.AccountPeriod.ID, &o.UserCreater.ID, &o.UserUpdater.ID, &o.Anulled, &o.CreatedAt, &o.UpdatedAt)
}
