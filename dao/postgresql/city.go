package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type CityDAOPsql struct {}

// Insert insertar registro en la BD
func (dao CityDAOPsql) Insert(obj *models.City) error {
	query := "INSERT INTO cities (department_id, code, city) VALUES ($1, upper($2), upper($3)) RETURNING id, department_id, code, city, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.DepartmentID, obj.Code, obj.City)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao CityDAOPsql) Update(obj *models.City) error {
	query := "UPDATE cities SET department_id = $2, code = upper($3), city = upper($4), updated_at = now() WHERE id = $1 RETURNING id, department_id, code, city, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.DepartmentID, obj.Code, obj.City)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao CityDAOPsql) Delete(obj *models.City) error {
	query := "DELETE FROM cities WHERE id = $1"
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
	obj = new(models.City)
	return nil
}

// GetByID consultar registro por id
func (dao CityDAOPsql) GetByID(id int) (*models.City, error) {
	query := "SELECT id, department_id, code, city, created_at, updated_at FROM cities WHERE id = $1"
	obj := &models.City{}
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
func (dao CityDAOPsql) GetAll() ([]models.City, error) {
	query := "SELECT id, department_id, code, city, created_at, updated_at FROM cities ORDER BY id"
	objs := make([]models.City, 0)
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
		var o models.City
		err = rows.Scan(&o.ID, &o.DepartmentID, &o.Code, &o.City, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao CityDAOPsql) rowToObject(row *sql.Row, o *models.City) error {
	return row.Scan(&o.ID, &o.DepartmentID, &o.Code, &o.City, &o.CreatedAt, &o.UpdatedAt)
}
