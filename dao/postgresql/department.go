package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type DepartmentDAOPsql struct{}

// Insert insertar registro en la BD
func (dao DepartmentDAOPsql) Insert(obj *models.Department) error {
	query := "INSERT INTO departments (code, department) VALUES (upper($1), upper($2)) RETURNING id, code, department, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Code, obj.Department)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao DepartmentDAOPsql) Update(obj *models.Department) error {
	query := "UPDATE departments SET code = upper($2), department = upper($3), updated_at = now() WHERE id = $1 RETURNING id, code, department, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Code, obj.Department)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao DepartmentDAOPsql) Delete(obj *models.Department) error {
	query := "DELETE FROM departments WHERE id = $1"
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
	obj = new(models.Department)
	return nil
}

// GetByID consultar registro por id
func (dao DepartmentDAOPsql) GetByID(id int) (*models.Department, error) {
	query := "SELECT id, code, department, created_at, updated_at FROM departments WHERE id = $1"
	obj := &models.Department{}
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
func (dao DepartmentDAOPsql) GetAll() ([]models.Department, error) {
	query := "SELECT id, code, department, created_at, updated_at FROM departments ORDER BY id"
	objs := make([]models.Department, 0)
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
		var o models.Department
		err = rows.Scan(&o.ID, &o.Code, &o.Department, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao DepartmentDAOPsql) GetCities(o *models.Department) error {
	query := "SELECT id, department_id, code, city, created_at, updated_at FROM cities WHERE department_id = $1"
	objs := make([]models.City, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(o.ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var obj models.City
		err = rows.Scan(&obj.ID, &obj.DepartmentID, &obj.Code, &obj.City, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return err
		}
		objs = append(objs, obj)
	}

	o.Cities = objs
	return nil
}

func (dao DepartmentDAOPsql) rowToObject(row *sql.Row, o *models.Department) error {
	return row.Scan(&o.ID, &o.Code, &o.Department, &o.CreatedAt, &o.UpdatedAt)
}
