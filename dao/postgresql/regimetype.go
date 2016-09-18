package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type RegimeTypeDAOPsql struct{}

// Insert insertar registro en la BD
func (dao RegimeTypeDAOPsql) Insert(obj *models.RegimeType) error {
	query := "INSERT INTO regime_type (regime) VALUES (upper($1)) RETURNING id, regime, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Regime)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao RegimeTypeDAOPsql) Update(obj *models.RegimeType) error {
	query := "UPDATE regime_type SET regime = upper($2), updated_at = now() WHERE id = $1 RETURNING id, regime, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Regime)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao RegimeTypeDAOPsql) Delete(obj *models.RegimeType) error {
	query := "DELETE FROM regime_type WHERE id = $1"
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
	obj = new(models.RegimeType)
	return nil
}

// GetByID consultar registro por id
func (dao RegimeTypeDAOPsql) GetByID(id int) (*models.RegimeType, error) {
	query := "SELECT id, regime, created_at, updated_at FROM regime_type WHERE id = $1"
	obj := &models.RegimeType{}
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
func (dao RegimeTypeDAOPsql) GetAll() ([]models.RegimeType, error) {
	query := "SELECT id, regime, created_at, updated_at FROM regime_type ORDER BY id"
	objs := make([]models.RegimeType, 0)
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
		var o models.RegimeType
		err = rows.Scan(&o.ID, &o.Regime, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao RegimeTypeDAOPsql) rowToObject(row *sql.Row, o *models.RegimeType) error {
	return row.Scan(&o.ID, &o.Regime, &o.CreatedAt, &o.UpdatedAt)
}
