package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type TaxDetailDAOPsql struct{}

// Insert insertar registro en la BD
func (dao TaxDetailDAOPsql) Insert(obj *models.TaxDetail) error {
	query := `INSERT INTO taxes_detail (tax_id, detail, percentage, account_puc_id, nature, base_value)
				VALUES ($1, upper($2), $3, upper($4), $5)
				RETURNING id, tax_id, detail, percentage, account_puc_id, nature, base_value, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Tax.ID, obj.Detail, obj.Percentage, obj.AccountPUC.ID, obj.Nature, obj.BaseValue)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao TaxDetailDAOPsql) Update(obj *models.TaxDetail) error {
	query := `UPDATE taxes_detail
				SET tax_id = $2, detail = upper($3), percentage = $4, account_puc_id = $5, nature = upper($6), base_value = $7, updated_at = now()
				WHERE id = $1
				RETURNING id, tax_id, detail, percentage, account_puc_id, nature, base_value, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Tax.ID, obj.Detail, obj.Percentage, obj.AccountPUC.ID, obj.Nature, obj.BaseValue)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao TaxDetailDAOPsql) Delete(obj *models.TaxDetail) error {
	query := "DELETE FROM taxes_detail WHERE id = $1"
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
	obj = new(models.TaxDetail)
	return nil
}

// GetByID consultar registro por id
func (dao TaxDetailDAOPsql) GetByID(id int) (*models.TaxDetail, error) {
	query := "SELECT id, tax_id, detail, percentage, account_puc_id, nature, base_value, created_at, updated_at FROM taxes_detail WHERE id = $1"
	obj := &models.TaxDetail{}
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
func (dao TaxDetailDAOPsql) GetAll() ([]models.TaxDetail, error) {
	query := "SELECT id, tax_id, detail, percentage, account_puc_id, nature, base_value, created_at, updated_at FROM taxes_detail ORDER BY id"
	objs := make([]models.TaxDetail, 0)
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
		var o models.TaxDetail
		err = rows.Scan(&o.ID, &o.Tax.ID, &o.Detail, &o.Percentage, &o.AccountPUC.ID, &o.Nature, &o.BaseValue, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao TaxDetailDAOPsql) rowToObject(row *sql.Row, o *models.TaxDetail) error {
	return row.Scan(&o.ID, &o.Tax.ID, &o.Detail, &o.Percentage, &o.AccountPUC.ID, &o.Nature, &o.BaseValue, &o.CreatedAt, &o.UpdatedAt)
}
