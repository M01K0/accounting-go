package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
)

type ObjectDAOPsql struct {}

// InsertObject insertar
func (o ObjectDAOPsql) InsertObject(object *models.Object) error {
	query := "INSERT INTO objetos (obj_codigo, obj_nombre, obj_descripcion) VALUES (upper($1), upper($2), $3) RETURNING obj_codigo, obj_nombre, obj_descripcion"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(object.Code, object.Name, object.Description).Scan(&object.Code, &object.Name, &object.Description)
	return err
}

// UpdateObject actualizar
func (o ObjectDAOPsql) UpdateObject(object *models.Object) error {
	query := "UPDATE objetos SET obj_codigo = upper($1), obj_nombre = upper($2), obj_descripcion = $3 WHERE obj_codigo = upper($1) RETURNING obj_codigo, obj_nombre, obj_descripcion"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(object.Code, object.Name, object.Description).Scan(&object.Code, &object.Name, &object.Description)
	return err
}

// DeleteObject borrar
func (o ObjectDAOPsql) DeleteObject(object *models.Object) error {
	query := "DELETE FROM objetos WHERE obj_codigo = upper($1)"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(object.Code)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("No se eliminó ningún registro")
	}
	return nil
}

// GetObjectByID Consulta por id
func (o ObjectDAOPsql) GetObjectByID(code string) (*models.Object, error) {
	query := "SELECT obj_codigo, obj_nombre, obj_descripcion FROM objetos WHERE obj_codigo = upper($1)"
	object := &models.Object{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(code).Scan(&object.Code, &object.Name, &object.Description)
	return object, err
}

// GetAllObject Consulta todos
func (o ObjectDAOPsql) GetAllObject() ([]models.Object, error) {
	query := "SELECT obj_codigo, obj_nombre, obj_descripcion FROM objetos"
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
		err = rows.Scan(&object.Code, &object.Name, &object.Description)
		if err != nil {
			return objects, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}
