package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type AccountPUCDAOPsql struct{}

// Insert insertar registro en la BD
func (dao AccountPUCDAOPsql) Insert(obj *models.AccountPUC) error {
	query := "SELECT * FROM fn_accounts_puc_insert($1, $2)"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Account, obj.AccountName)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao AccountPUCDAOPsql) Update(obj *models.AccountPUC) error {
	query := "UPDATE accounts_puc SET account_name = upper($2), updated_at = now() WHERE id = $1 RETURNING id, account, account_name, account_puc_parent_id, account_class_id, account_level_id, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.AccountName)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao AccountPUCDAOPsql) Delete(obj *models.AccountPUC) error {
	query := "DELETE FROM accounts_puc WHERE id = $1"
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
	obj = new(models.AccountPUC)
	return nil
}

// GetByID consultar registro por id
func (dao AccountPUCDAOPsql) GetByID(id int) (*models.AccountPUC, error) {
	query := "SELECT id, account, account_name, account_puc_parent_id, account_class_id, account_level_id, created_at, updated_at FROM accounts_puc WHERE id = $1"
	obj := &models.AccountPUC{}
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

func (dao AccountPUCDAOPsql) GetByAccount(a string) (*models.AccountPUC, error) {
	query := "SELECT id, account, account_name, account_puc_parent_id, account_class_id, account_level_id, created_at, updated_at FROM accounts_puc WHERE account = $1"
	obj := &models.AccountPUC{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(a)
	err = dao.rowToObject(row, obj)
	return obj, err
}

// GetAll Consulta todos los registros de la bd
func (dao AccountPUCDAOPsql) GetAll() ([]models.AccountPUC, error) {
	query := "SELECT id, account, account_name, account_puc_parent_id, account_class_id, account_level_id, created_at, updated_at FROM accounts_puc ORDER BY id"
	objs := make([]models.AccountPUC, 0)
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
		var o models.AccountPUC
		err = rows.Scan(&o.ID, &o.Account, &o.AccountName, &o.AccountParentID, &o.AccountClass.ID, &o.AccountClass.ID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountPUCDAOPsql) rowToObject(row *sql.Row, o *models.AccountPUC) error {
	return row.Scan(&o.ID, &o.Account, &o.AccountName, &o.AccountParentID, &o.AccountClass.ID, &o.AccountClass.ID, &o.CreatedAt, &o.UpdatedAt)
}
