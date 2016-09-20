package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type BalanceSheetHeaderDAOPsql struct {}

// Insert insertar registro en la BD
func (dao BalanceSheetHeaderDAOPsql) Insert(obj *models.BalanceSheetHeader) error {
	query := "INSERT INTO balance_sheet_header (user_id, account_period_id) VALUES ($1, $2) RETURNING id, user_id, account_period_id, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.User.ID, obj.AccountPeriod.ID)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao BalanceSheetHeaderDAOPsql) Update(obj *models.BalanceSheetHeader) error {
	query := "UPDATE balance_sheet_header SET user_id = $2, account_period_id = $3, updated_at = now() WHERE id = $1 RETURNING id, user_id, account_period_id, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.User.ID, obj.AccountPeriod.ID)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao BalanceSheetHeaderDAOPsql) Delete(obj *models.BalanceSheetHeader) error {
	query := "DELETE FROM balance_sheet_header WHERE id = $1"
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
	obj = new(models.BalanceSheetHeader)
	return nil
}

// GetByID consultar registro por id
func (dao BalanceSheetHeaderDAOPsql) GetByID(id int) (*models.BalanceSheetHeader, error) {
	query := "SELECT id, user_id, account_period_id, created_at, updated_at FROM balance_sheet_header WHERE id = $1"
	obj := &models.BalanceSheetHeader{}
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
func (dao BalanceSheetHeaderDAOPsql) GetAll() ([]models.BalanceSheetHeader, error) {
	query := "SELECT id, user_id, account_period_id, created_at, updated_at FROM balance_sheet_header ORDER BY id"
	objs := make([]models.BalanceSheetHeader, 0)
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
		var o models.BalanceSheetHeader
		err = rows.Scan(&o.ID, &o.User.ID, &o.AccountPeriod.ID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao BalanceSheetHeaderDAOPsql) rowToObject(row *sql.Row, o *models.BalanceSheetHeader) error {
	return row.Scan(&o.ID, &o.User.ID, &o.AccountPeriod.ID, &o.CreatedAt, &o.UpdatedAt)
}
