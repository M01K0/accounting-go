package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type CostCenterDAOPsql struct{}

// Insert insertar registro en la BD
func (dao CostCenterDAOPsql) Insert(obj *models.CostCenter) error {
	query := "INSERT INTO cost_centers (code, cost_center) VALUES (upper($1), upper($2)) RETURNING id, code, cost_center, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Code, obj.CostCenter)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao CostCenterDAOPsql) Update(obj *models.CostCenter) error {
	query := "UPDATE cost_centers SET code = upper($2), cost_center = upper($3), updated_at = now() WHERE id = $1 RETURNING id, code, cost_center, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Code, obj.CostCenter)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao CostCenterDAOPsql) Delete(obj *models.CostCenter) error {
	query := "DELETE FROM cost_centers WHERE id = $1"
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
	obj = new(models.CostCenter)
	return nil
}

// GetByID consultar registro por id
func (dao CostCenterDAOPsql) GetByID(id int) (*models.CostCenter, error) {
	query := "SELECT id, code, cost_center, created_at, updated_at FROM cost_centers WHERE id = $1"
	obj := &models.CostCenter{}
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
func (dao CostCenterDAOPsql) GetAll() ([]models.CostCenter, error) {
	query := "SELECT id, code, cost_center, created_at, updated_at FROM cost_centers ORDER BY id"
	objs := make([]models.CostCenter, 0)
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
		var o models.CostCenter
		err = rows.Scan(&o.ID, &o.Code, &o.CostCenter, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao CostCenterDAOPsql) rowToObject(row *sql.Row, o *models.CostCenter) error {
	return row.Scan(&o.ID, &o.Code, &o.CostCenter, &o.CreatedAt, &o.UpdatedAt)
}
