package postgresql

import "github.com/alexyslozada/accounting-go/models"

// ProfileInsert inserta un registro a la BD
func ProfileInsert(profile *models.Profile) error {
	query := "INSERT INTO perfiles (nombre) VALUES ($1) RETURNING idperfil, nombre, activo"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(profile.Name).Scan(&profile.ID, &profile.Name, &profile.Active)

	return nil
}

/*
// Update Actualiza el registro en la BD
func (profile *models.Profile) Update() error {
	query := "SELECT fn_perfiles_upd($1,$2,$3)"
	response := false

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(profile.ID, profile.Name, profile.Active).Scan(&response)
	if err != nil {
		return err
	}

	return nil
}

// Delete Borra un registro de la BD
func (profile *models.Profile) Delete() error {
	query := "SELECT fn_periles_del($1)"
	response := false

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(profile.ID).Scan(&response)
	if err != nil {
		return err
	}

	return nil
}

// Get obtiene un registro de la BD
func (profile *models.Profile) Get() error {
	query := "SELECT idperfil, nombre, activo FROM perfiles WHERE idperfil = $1"

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(profile.ID).Scan(profile.ID, profile.Name, profile.Active)
	if err != nil {
		return err
	}

	return nil
}
*/
