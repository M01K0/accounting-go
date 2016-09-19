package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type AccountPeriodDAOPsql struct {}

// Insert insertar registro en la BD
func (dao AccountPeriodDAOPsql) Insert(obj *models.AccountPeriod) error {
	query := "INSERT INTO account_period (year, month) VALUES ($1, $2) RETURNING id, year, month, open, close_date, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Year, obj.Month)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao AccountPeriodDAOPsql) Update(obj *models.AccountPeriod) error {
	query := "UPDATE account_period SET year = $2, month = $3, open = $4, close_date = $5, updated_at = now() WHERE id = $1 RETURNING id, year, month, open, close_date, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Year, obj.Month, obj.Open, obj.CloseDate)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao AccountPeriodDAOPsql) Delete(obj *models.AccountPeriod) error {
	query := "DELETE FROM account_period WHERE id = $1"
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
	obj = new(models.AccountPeriod)
	return nil
}

// GetByID consultar registro por id
func (dao AccountPeriodDAOPsql) GetByID(id int) (*models.AccountPeriod, error) {
	query := "SELECT id, year, month, open, close_date, created_at, updated_at FROM account_period WHERE id = $1"
	obj := &models.AccountPeriod{}
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
func (dao AccountPeriodDAOPsql) GetAll() ([]models.AccountPeriod, error) {
	query := "SELECT id, year, month, open, close_date, created_at, updated_at FROM AccountPeriod ORDER BY id"
	objs := make([]models.AccountPeriod, 0)
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
		var o models.AccountPeriod
		err = rows.Scan(&o.ID, &o.Year, &o.Month, &o.Open, &o.CloseDate, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountPeriodDAOPsql) rowToObject(row *sql.Row, o *models.AccountPeriod) error {
	return row.Scan(&o.ID, &o.Year, &o.Month, &o.Open, &o.CloseDate, &o.CreatedAt, &o.UpdatedAt)
}
