package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type TaxpayerTypeDAOPsql struct{}

// Insert insertar registro en la BD
func (dao TaxpayerTypeDAOPsql) Insert(obj *models.TaxpayerType) error {
	query := "INSERT INTO taxpayer_type (taxpayer) VALUES (upper($1)) RETURNING id, taxpayer, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Taxpayer)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao TaxpayerTypeDAOPsql) Update(obj *models.TaxpayerType) error {
	query := "UPDATE taxpayer_type SET taxpayer = upper($2), updated_at = now() WHERE id = $1 RETURNING id, taxpayer, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Taxpayer)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao TaxpayerTypeDAOPsql) Delete(obj *models.TaxpayerType) error {
	query := "DELETE FROM taxpayer_type WHERE id = $1"
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
	obj = new(models.TaxpayerType)
	return nil
}

// GetByID consultar registro por id
func (dao TaxpayerTypeDAOPsql) GetByID(id int) (*models.TaxpayerType, error) {
	query := "SELECT id, taxpayer, created_at, updated_at FROM taxpayer_type WHERE id = $1"
	obj := &models.TaxpayerType{}
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
func (dao TaxpayerTypeDAOPsql) GetAll() ([]models.TaxpayerType, error) {
	query := "SELECT id, taxpayer, created_at, updated_at FROM taxpayer_type ORDER BY id"
	objs := make([]models.TaxpayerType, 0)
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
		var o models.TaxpayerType
		err = rows.Scan(&o.ID, &o.Taxpayer, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao TaxpayerTypeDAOPsql) rowToObject(row *sql.Row, o *models.TaxpayerType) error {
	return row.Scan(&o.ID, &o.Taxpayer, &o.CreatedAt, &o.UpdatedAt)
}
