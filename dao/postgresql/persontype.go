package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type PersonTypeDAOPsql struct{}

// Insert insertar registro en la BD
func (dao PersonTypeDAOPsql) Insert(obj *models.PersonType) error {
	query := "INSERT INTO person_type (person) VALUES (upper($1)) RETURNING id, person, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Person)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao PersonTypeDAOPsql) Update(obj *models.PersonType) error {
	query := "UPDATE person_type SET person = upper($2), updated_at = now() WHERE id = $1 RETURNING id, person, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Person)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao PersonTypeDAOPsql) Delete(obj *models.PersonType) error {
	query := "DELETE FROM person_type WHERE id = $1"
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
	obj = new(models.PersonType)
	return nil
}

// GetByID consultar registro por id
func (dao PersonTypeDAOPsql) GetByID(id int) (*models.PersonType, error) {
	query := "SELECT id, person created_at, updated_at FROM person_type WHERE id = $1"
	obj := &models.PersonType{}
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
func (dao PersonTypeDAOPsql) GetAll() ([]models.PersonType, error) {
	query := "SELECT id, person, created_at, updated_at FROM person_type ORDER BY id"
	objs := make([]models.PersonType, 0)
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
		var o models.PersonType
		err = rows.Scan(&o.ID, &o.Person, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao PersonTypeDAOPsql) rowToObject(row *sql.Row, o *models.PersonType) error {
	return row.Scan(&o.ID, &o.Person, &o.CreatedAt, &o.UpdatedAt)
}
