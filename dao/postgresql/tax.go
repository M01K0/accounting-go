package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type TaxDAOPsql struct{}

// Insert insertar registro en la BD
func (dao TaxDAOPsql) Insert(obj *models.Tax) error {
	query := "INSERT INTO taxes (tax) VALUES (upper($1)) RETURNING id, tax, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Tax)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao TaxDAOPsql) Update(obj *models.Tax) error {
	query := "UPDATE taxes SET tax = upper($2), updated_at = now() WHERE id = $1 RETURNING id, tax, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Tax)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao TaxDAOPsql) Delete(obj *models.Tax) error {
	query := "DELETE FROM taxes WHERE id = $1"
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
	obj = new(models.Tax)
	return nil
}

// GetByID consultar registro por id
func (dao TaxDAOPsql) GetByID(id int) (*models.Tax, error) {
	query := "SELECT id, tax, created_at, updated_at FROM taxes WHERE id = $1"
	obj := &models.Tax{}
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
func (dao TaxDAOPsql) GetAll() ([]models.Tax, error) {
	query := "SELECT id, tax, created_at, updated_at FROM taxes ORDER BY id"
	objs := make([]models.Tax, 0)
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
		var o models.Tax
		err = rows.Scan(&o.ID, &o.Tax, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao TaxDAOPsql) rowToObject(row *sql.Row, o *models.Tax) error {
	return row.Scan(&o.ID, &o.Tax, &o.CreatedAt, &o.UpdatedAt)
}
