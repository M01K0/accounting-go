package postgresql

import (
	"database/sql"
	"github.com/alexyslozada/accounting-go/models"
)

type BalanceSheetDetailDAOPsql struct{}

// Insert insertar registro en la BD
func (dao BalanceSheetDetailDAOPsql) Insert(obj *models.BalanceSheetDetail) error {
	query := `INSERT INTO balance_sheet_detail (balance_sheet_header_id, account_puc_id, third_party_id, cost_center_id, previous_balance, debit, credit, current_balance)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
				RETURNING id, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.BalanceSheetHeaderID, obj.AccountPUC.ID, obj.ThirdParty.ID, obj.CostCenter.ID, obj.PreviousBalance, obj.Debit, obj.Credit, obj.CurrentBalance)
	return dao.rowToObject(row, obj)
}

// GetByID consultar registro por id
func (dao BalanceSheetDetailDAOPsql) GetByID(id int) (*models.BalanceSheetDetail, error) {
	query := "SELECT id, created_at, updated_at FROM BalanceSheetDetail WHERE id = $1"
	obj := &models.BalanceSheetDetail{}
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
func (dao BalanceSheetDetailDAOPsql) GetAll() ([]models.BalanceSheetDetail, error) {
	query := "SELECT id, created_at, updated_at FROM BalanceSheetDetail ORDER BY id"
	objs := make([]models.BalanceSheetDetail, 0)
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
		var o models.BalanceSheetDetail
		err = rows.Scan(&o.ID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao BalanceSheetDetailDAOPsql) rowToObject(row *sql.Row, o *models.BalanceSheetDetail) error {
	return row.Scan(&o.ID, &o.CreatedAt, &o.UpdatedAt)
}
