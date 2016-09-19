package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type AccountLevelDAOPsql struct {}

// Insert insertar registro en la BD
func (dao AccountLevelDAOPsql) Insert(obj *models.AccountLevel) error {
	query := "INSERT INTO account_levels (account_level, digits) VALUES (upper($1), $2) RETURNING id, account_level, digits, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.AccountLevel, obj.Digits)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao AccountLevelDAOPsql) Update(obj *models.AccountLevel) error {
	query := "UPDATE account_levels SET account_level = upper($2), digits = $3, updated_at = now() WHERE id = $1 RETURNING id, account_level, digits, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.AccountLevel, obj.Digits)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao AccountLevelDAOPsql) Delete(obj *models.AccountLevel) error {
	query := "DELETE FROM account_levels WHERE id = $1"
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
	obj = new(models.AccountLevel)
	return nil
}

// GetByID consultar registro por id
func (dao AccountLevelDAOPsql) GetByID(id int) (*models.AccountLevel, error) {
	query := "SELECT id, account_level, digits, created_at, updated_at FROM account_levels WHERE id = $1"
	obj := &models.AccountLevel{}
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
func (dao AccountLevelDAOPsql) GetAll() ([]models.AccountLevel, error) {
	query := "SELECT id, account_level, digits, created_at, updated_at FROM account_levels ORDER BY id"
	objs := make([]models.AccountLevel, 0)
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
		var o models.AccountLevel
		err = rows.Scan(&o.ID, &o.AccountLevel, &o.Digits, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountLevelDAOPsql) rowToObject(row *sql.Row, o *models.AccountLevel) error {
	return row.Scan(&o.ID, &o.AccountLevel, &o.Digits, &o.CreatedAt, &o.UpdatedAt)
}
