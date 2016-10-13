package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type AccountingDetailDAOPsql struct{}

// Insert insertar registro en la BD
func (dao AccountingDetailDAOPsql) Insert(obj *models.AccountingDetail) error {
	query := `INSERT INTO accounting_details (accounting_header_id, account_puc_id, debit, credit, third_party_id, cost_center_id)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id, accounting_header_id, account_puc_id, debit, credit, third_party_id, cost_center_id, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.AccountingHeaderID, obj.AccountPUC.ID, obj.Debit, obj.Credit, obj.ThirdParty.ID, obj.CostCenter.ID)
	return dao.rowToObject(row, obj)
}

// DeleteByHeader borrar registro de la bd por el número de id del encabezado
func (dao AccountingDetailDAOPsql) DeleteByHeader(idh int) error {
	query := "DELETE FROM accounting_details WHERE accounting_header_id = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	result, err := stmt.Exec(idh)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("No se eliminó ningún registro")
	}
	return nil
}

// GetByID consultar registro por id
func (dao AccountingDetailDAOPsql) GetByID(id int) (*models.AccountingDetail, error) {
	query := "SELECT id, accounting_header_id, account_puc_id, debit, credit, third_party_id, cost_center_id, created_at, updated_at FROM accounting_details WHERE id = $1"
	obj := &models.AccountingDetail{}
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
func (dao AccountingDetailDAOPsql) GetAll() ([]models.AccountingDetail, error) {
	query := "SELECT id, accounting_header_id, account_puc_id, debit, credit, third_party_id, cost_center_id, created_at, updated_at FROM accounting_details ORDER BY id"
	objs := make([]models.AccountingDetail, 0)
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
		var o models.AccountingDetail
		err = rows.Scan(&o.ID, &o.AccountingHeaderID, &o.AccountPUC.ID, &o.Debit, &o.Credit, &o.ThirdParty.ID, &o.CostCenter.ID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountingDetailDAOPsql) rowToObject(row *sql.Row, o *models.AccountingDetail) error {
	return row.Scan(&o.ID, &o.AccountingHeaderID, &o.AccountPUC.ID, &o.Debit, &o.Credit, &o.ThirdParty.ID, &o.CostCenter.ID, &o.CreatedAt, &o.UpdatedAt)
}
