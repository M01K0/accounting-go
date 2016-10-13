package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type FunctionaryCompanyDAOPsql struct{}

// Insert insertar registro en la BD
func (dao FunctionaryCompanyDAOPsql) Insert(obj *models.FunctionaryCompany) error {
	query := `INSERT INTO functionary_company (functionary_type_id, identification_type_id, identification_number, verification_digit, functinoary)
				VALUES ($1, $2, $3, $4, upper($5), $6)
				RETURNING id, functionary_type_id, identification_type_id, identification_number, verification_digit, functionary, active, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.FunctionaryType.ID, obj.IdentificationType.ID, obj.IdentificationNumber, obj.VerificationDigit, obj.Functionary)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao FunctionaryCompanyDAOPsql) Update(obj *models.FunctionaryCompany) error {
	query := `UPDATE functionary_company SET functionary_type_id = $2, identification_type_id = $3, identification_number = $4, verification_digit = $5, functionary = upper($6), active = $7, updated_at = now()
		"		WHERE id = $1 RETURNING id, functionary_type_id, identification_type_id, identification_number, verification_digit, functionary, active, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.FunctionaryType.ID, obj.IdentificationType.ID, obj.IdentificationNumber, obj.VerificationDigit, obj.Functionary, obj.Active)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao FunctionaryCompanyDAOPsql) Delete(obj *models.FunctionaryCompany) error {
	query := "DELETE FROM functionary_company WHERE id = $1"
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
	obj = new(models.FunctionaryCompany)
	return nil
}

// GetByID consultar registro por id
func (dao FunctionaryCompanyDAOPsql) GetByID(id int) (*models.FunctionaryCompany, error) {
	query := `SELECT id, functionary_type_id, identification_type_id, identification_number, verification_digit, functionary, active, created_at, updated_at
				FROM functionary_company WHERE id = $1`
	obj := &models.FunctionaryCompany{}
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
func (dao FunctionaryCompanyDAOPsql) GetAll() ([]models.FunctionaryCompany, error) {
	query := `SELECT id, functionary_type_id, identification_type_id, identification_number, verification_digit, functionary, active, created_at, updated_at
				FROM functionary_company ORDER BY id`
	objs := make([]models.FunctionaryCompany, 0)
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
		var o models.FunctionaryCompany
		err = rows.Scan(&o.ID, &o.FunctionaryType.ID, &o.IdentificationType.ID, &o.IdentificationNumber, &o.VerificationDigit, &o.Functionary, &o.Active, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao FunctionaryCompanyDAOPsql) rowToObject(row *sql.Row, o *models.FunctionaryCompany) error {
	return row.Scan(&o.ID, &o.FunctionaryType.ID, &o.IdentificationType.ID, &o.IdentificationNumber, &o.VerificationDigit, &o.Functionary, &o.Active, &o.CreatedAt, &o.UpdatedAt)
}
