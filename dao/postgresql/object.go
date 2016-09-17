package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

// ObjectDAOPsql estructura dao de object
type ObjectDAOPsql struct {}

// Insert insertar
func (dao ObjectDAOPsql) Insert(object *models.Object) error {
	query := "INSERT INTO objects (code, object_name, description) VALUES (upper($1), upper($2), $3) RETURNING id, code, object_name, description, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(object.Code, object.ObjectName, object.Description)
	return dao.rowToObject(row, object)
}

// Update actualizar
func (dao ObjectDAOPsql) Update(object *models.Object) error {
	query := "UPDATE objects SET code = upper($1), object_name = upper($2), description = upper($3), updated_at = now() WHERE id = $4 RETURNING id, code, object_name, description, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(object.Code, object.ObjectName, object.Description, object.ID)
	return dao.rowToObject(row, object)

}

// Delete borrar
func (dao ObjectDAOPsql) Delete(object *models.Object) error {
	query := "DELETE FROM objects WHERE id = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(object.ID)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("No se eliminó ningún registro")
	}
	object = new(models.Object)
	return nil
}

// GetByID Consulta por id
func (dao ObjectDAOPsql) GetByID(id int) (*models.Object, error) {
	query := "SELECT id, code, object_name, description, created_at, updated_at FROM objects WHERE id = $1"
	object := &models.Object{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = dao.rowToObject(row, object)
	return object, err
}

// GetAll Consulta todos
func (dao ObjectDAOPsql) GetAll() ([]models.Object, error) {
	query := "SELECT id, code, object_name, description, created_at, updated_at FROM objects ORDER BY code"
	objects := make([]models.Object, 0)
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
		var object models.Object
		err = rows.Scan(&object.ID, &object.Code, &object.ObjectName, &object.Description, &object.CreatedAt, &object.UpdatedAt)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}

// rowToObject mapea la consulta al objeto
func (dao ObjectDAOPsql) rowToObject(row *sql.Row, object *models.Object) error {
	return row.Scan(&object.ID, &object.Code, &object.ObjectName, &object.Description, &object.CreatedAt, &object.UpdatedAt)
}
