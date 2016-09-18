package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type FunctionaryTypeDAOPsql struct {}

// Insert insertar registro en la BD
func (dao FunctionaryTypeDAOPsql) Insert(obj *models.FunctionaryType) error {
	query := "INSERT INTO functionary_type (functionary) VALUES (upper($1)) RETURNING id, functionary, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Functionary)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao FunctionaryTypeDAOPsql) Update(obj *models.FunctionaryType) error {
	query := "UPDATE functionary_type SET functionary = upper($2), updated_at = now() WHERE id = $1 RETURNING id, functionary, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Functionary)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao FunctionaryTypeDAOPsql) Delete(obj *models.FunctionaryType) error {
	query := "DELETE FROM functionary_type WHERE id = $1"
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
	obj = new(models.FunctionaryType)
	return nil
}

// GetByID consultar registro por id
func (dao FunctionaryTypeDAOPsql) GetByID(id int) (*models.FunctionaryType, error) {
	query := "SELECT id, functionary, created_at, updated_at FROM functionary_type WHERE id = $1"
	obj := &models.FunctionaryType{}
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
func (dao FunctionaryTypeDAOPsql) GetAll() ([]models.FunctionaryType, error) {
	query := "SELECT id, functionary, created_at, updated_at FROM FunctionaryType ORDER BY id"
	objs := make([]models.FunctionaryType, 0)
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
		var o models.FunctionaryType
		err = rows.Scan(&o.ID, &o.Functionary, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao FunctionaryTypeDAOPsql) rowToObject(row *sql.Row, o *models.FunctionaryType) error {
	return row.Scan(&o.ID, &o.Functionary, &o.CreatedAt, &o.UpdatedAt)
}
