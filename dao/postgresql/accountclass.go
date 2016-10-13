package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type AccountClassDAOPsql struct{}

// Insert insertar registro en la BD
func (dao AccountClassDAOPsql) Insert(obj *models.AccountClass) error {
	query := "INSERT INTO account_classes (report_type_id, account_class, nature) VALUES ($1, upper($2), upper($3)) RETURNING id, report_type_id, account_class, nature, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ReportType.ID, obj.AccountClass, obj.Nature)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao AccountClassDAOPsql) Update(obj *models.AccountClass) error {
	query := "UPDATE account_classes SET report_type_id = $2, account_class = upper($3), nature = upper($4), updated_at = now() WHERE id = $1 RETURNING id, report_type_id, account_class, nature, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.ReportType.ID, obj.AccountClass, obj.Nature)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao AccountClassDAOPsql) Delete(obj *models.AccountClass) error {
	query := "DELETE FROM account_classes WHERE id = $1"
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
	obj = new(models.AccountClass)
	return nil
}

// GetByID consultar registro por id
func (dao AccountClassDAOPsql) GetByID(id int) (*models.AccountClass, error) {
	query := "SELECT id, report_type_id, acount_class, nature, created_at, updated_at FROM account_class WHERE id = $1"
	obj := &models.AccountClass{}
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
func (dao AccountClassDAOPsql) GetAll() ([]models.AccountClass, error) {
	query := "SELECT id, report_type_id, acount_class, nature, created_at, updated_at FROM account_classes ORDER BY id"
	objs := make([]models.AccountClass, 0)
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
		var o models.AccountClass
		err = rows.Scan(&o.ID, &o.ReportType.ID, &o.AccountClass, &o.Nature, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao AccountClassDAOPsql) rowToObject(row *sql.Row, o *models.AccountClass) error {
	return row.Scan(&o.ID, &o.ReportType.ID, &o.AccountClass, &o.Nature, &o.CreatedAt, &o.UpdatedAt)
}
