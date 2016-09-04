package postgresql

import (
	"errors"
	"github.com/alexyslozada/accounting-go/models"
	"database/sql"
)

type ProfileDAOPsql struct{}

// ProfileInsert inserta un registro a la BD
func (dao ProfileDAOPsql) InsertProfile(profile *models.Profile) error {
	query := "INSERT INTO perfiles (nombre) VALUES (upper($1)) RETURNING idperfil, nombre, activo"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(profile.Name).Scan(&profile.ID, &profile.Name, &profile.Active)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProfile Actualiza el registro en la BD
func (dao ProfileDAOPsql) UpdateProfile(profile *models.Profile) error {
	query := "UPDATE perfiles SET nombre = upper($1), activo = $2 WHERE idperfil = $3 RETURNING idperfil, nombre, activo"

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(profile.Name, profile.Active, profile.ID).Scan(&profile.ID, &profile.Name, &profile.Active)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProfile Borra un registro de la BD
func (dao ProfileDAOPsql) DeleteProfile(profile *models.Profile) error {
	query := "DELETE FROM perfiles WHERE idperfil = $1"

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
	if rowsaffected, _ := result.RowsAffected(); rowsaffected == 0 {
		return errors.New("No se eliminó ningún registro")
	}
	return nil
}

// GetProfileByID obtiene un registro de la BD
func (dao ProfileDAOPsql) GetProfileByID(id int16) (*models.Profile, error) {
	query := "SELECT idperfil, nombre, activo FROM perfiles WHERE idperfil = $1"
	profile := &models.Profile{}

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&profile.ID, &profile.Name, &profile.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No se encontraron registros")
		}
		return nil, err
	}

	return profile, nil
}

// GetAllProfiles obtiene todos los perfiles de la BD
func (dao ProfileDAOPsql) GetAllProfiles() ([]models.Profile, error) {
	query := "SELECT idperfil, nombre, activo FROM perfiles"
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
		err = rows.Scan(&profile.ID, &profile.Name, &profile.Active)
		if err != nil {
			return profiles, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
