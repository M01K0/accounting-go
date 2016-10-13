package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

// ProfileDAOPsql estructura dao de profile
type ProfileDAOPsql struct{}

// Insert inserta un registro a la BD
func (dao ProfileDAOPsql) Insert(profile *models.Profile) error {
	query := "INSERT INTO profiles (profile) VALUES (upper($1)) RETURNING id, profile, active, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(profile.Profile)
	return dao.rowToObject(row, profile)
}

// Update Actualiza el registro en la BD
func (dao ProfileDAOPsql) Update(profile *models.Profile) error {
	query := "UPDATE profiles SET profile = upper($1), active = $2, updated_at = now() WHERE id = $3 RETURNING id, profile, active, created_at, updated_at"

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(profile.Profile, profile.Active, profile.ID)
	return dao.rowToObject(row, profile)
}

// Delete Borra un registro de la BD
func (dao ProfileDAOPsql) Delete(profile *models.Profile) error {
	query := "DELETE FROM profiles WHERE id = $1"

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(profile.ID)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("No se eliminó ningún registro")
	}
	profile = new(models.Profile)
	return nil
}

// GetByID obtiene un registro de la BD
func (dao ProfileDAOPsql) GetByID(id int16) (*models.Profile, error) {
	query := "SELECT id, profile, active, created_at, updated_at FROM profiles WHERE id = $1"
	profile := &models.Profile{}

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = dao.rowToObject(row, profile)
	return profile, err
}

// GetAll obtiene todos los perfiles de la BD
func (dao ProfileDAOPsql) GetAll() ([]models.Profile, error) {
	query := "SELECT id, profile, active, created_at, updated_at FROM profiles ORDER BY id"
	profiles := make([]models.Profile, 0)

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
		var profile models.Profile
		err = rows.Scan(&profile.ID, &profile.Profile, &profile.Active, &profile.CreatedAt, &profile.UpdatedAt)
		if err != nil {
			return profiles, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

// rowToObject permite mapear la consulta al objeto
func (dao ProfileDAOPsql) rowToObject(row *sql.Row, object *models.Profile) error {
	return row.Scan(&object.ID, &object.Profile, &object.Active, &object.CreatedAt, &object.UpdatedAt)
}
