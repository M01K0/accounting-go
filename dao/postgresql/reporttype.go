package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type ReportTypeDAOPsql struct{}

// Insert insertar registro en la BD
func (dao ReportTypeDAOPsql) Insert(obj *models.ReportType) error {
	query := "INSERT INTO report_type (report) VALUES (upper($1)) RETURNING id, report, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Report)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao ReportTypeDAOPsql) Update(obj *models.ReportType) error {
	query := "UPDATE report_type SET report = upper($2), updated_at = now() WHERE id = $1 RETURNING id, report, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Report)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao ReportTypeDAOPsql) Delete(obj *models.ReportType) error {
	query := "DELETE FROM report_type WHERE id = $1"
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
	obj = new(models.ReportType)
	return nil
}

// GetByID consultar registro por id
func (dao ReportTypeDAOPsql) GetByID(id int) (*models.ReportType, error) {
	query := "SELECT id, report, created_at, updated_at FROM report_type WHERE id = $1"
	obj := &models.ReportType{}
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
func (dao ReportTypeDAOPsql) GetAll() ([]models.ReportType, error) {
	query := "SELECT id, report, created_at, updated_at FROM report_type ORDER BY id"
	objs := make([]models.ReportType, 0)
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
		var o models.ReportType
		err = rows.Scan(&o.ID, &o.Report, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao ReportTypeDAOPsql) rowToObject(row *sql.Row, o *models.ReportType) error {
	return row.Scan(&o.ID, &o.Report, &o.CreatedAt, &o.UpdatedAt)
}
