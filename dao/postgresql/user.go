package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type UserDAOPsql struct{}

// Insert insertar registro en la BD
func (dao UserDAOPsql) Insert(obj *models.User) error {
	query := "INSERT INTO users (identification, username, email, passwd, profile_id) VALUES ($1, upper($2), $3, md5($4), $5) RETURNING id, identification, username, email, passwd, profile_id, active, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao UserDAOPsql) Update(obj *models.User) error {
	query := "UPDATE users SET identification = $2, username = upper($3), email = $4, passwd = $5, profile_id = $6, active = $7, updated_at = now() WHERE id = $1 RETURNING id, identification, username, email, passwd, profile_id, active, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Identification, obj.Username, obj.Email, obj.Passwd, obj.Profile.ID, obj.Active)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao UserDAOPsql) Delete(obj *models.User) error {
	query := "DELETE FROM users WHERE id = $1"
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
	obj = new(models.User)
	return nil
}

// GetByID consultar registro por id
func (dao UserDAOPsql) GetByID(id int) (*models.User, error) {
	query := "SELECT id, identification, username, email, passwd, profile_id, active, created_at, updated_at FROM users WHERE id = $1"
	obj := &models.User{}
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
func (dao UserDAOPsql) GetAll() ([]models.User, error) {
	query := "SELECT id, identification, username, email, passwd, profile_id, active, created_at, updated_at FROM users ORDER BY id"
	objs := make([]models.User, 0)
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
		var o models.User
		err = rows.Scan(&o.ID, &o.ID, &o.Identification, &o.Username, &o.Email, &o.Passwd, &o.Profile.ID, &o.Active, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao UserDAOPsql) rowToObject(row *sql.Row, o *models.User) error {
	return row.Scan(&o.ID, &o.Identification, &o.Username, &o.Email, &o.Passwd, &o.Profile.ID, &o.Active, &o.CreatedAt, &o.UpdatedAt)
}
