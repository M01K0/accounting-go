package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

type CompanyDAOPsql struct {}

// Insert insertar registro en la BD
func (dao CompanyDAOPsql) Insert(obj *models.Company) error {
	query := `INSERT INTO companies (identification_type, identification_number, verification_digit, company, address, phone, departments_id, cities_id, web, email, activity, autorretenedor, person_type_id, regime_type_id, taxpayer_type_id, logo)
				VALUES ($1, $2, $3, upper($4), upper($5), $6, $7, $8, lower($9), lower($10), $11, $12, $13, $14, $15, lower($16))
				RETURNING id, identification_type_id, identification_number, verification_digit, company, address, phone, departments_id, cities_id, web, email, activity, autorretenedor, person_type_id, regime_type_id, taxpayer_type_id, logo, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.IdentificationType.ID, obj.IdentificationNumber, obj.VerificationDigit, obj.Company, obj.Address, obj.Phone, obj.Department.ID, obj.City.ID, obj.Web, obj.Email, obj.Activity, obj.AutoRretenedor, obj.PersonType.ID, obj.RegimeType.ID, obj.TaxpayerType.ID, obj.Logo)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao CompanyDAOPsql) Update(obj *models.Company) error {
	query := `UPDATE companies SET identification_type_id = $2, identification_number = $3, verification_digit = $4, company = upper($5), address = upper($6), phone = $7, department_id = $8, city_id = $9, web = lower($10), email = lower($11), activity = $12, autorretenedor = $13, person_type_id = $14, regime_type_id = $15, taxpayer_type_id = $16, logo = lower($17), updated_at = now() WHERE id = $1
				RETURNING id, identification_type_id, identification_number, verification_digit, company, address, phone, departments_id, cities_id, web, email, activity, autorretenedor, person_type_id, regime_type_id, taxpayer_type_id, logo, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.IdentificationType.ID, obj.IdentificationNumber, obj.VerificationDigit, obj.Company, obj.Address, obj.Phone, obj.Department.ID, obj.City.ID, obj.Web, obj.Email, obj.Activity, obj.AutoRretenedor, obj.PersonType.ID, obj.RegimeType.ID, obj.TaxpayerType.ID, obj.Logo)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao CompanyDAOPsql) Delete(obj *models.Company) error {
	query := "DELETE FROM companies WHERE id = $1"
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
	obj = new(models.Company)
	return nil
}

// GetByID consultar registro por id
func (dao CompanyDAOPsql) GetByID(id int) (*models.Company, error) {
	query := `SELECT id, identification_type_id, identification_number, verification_digit, company, address, phone, departments_id, cities_id, web, email, activity, autorretenedor, person_type_id, regime_type_id, taxpayer_type_id, logo, created_at, updated_at
				FROM companies WHERE id = $1`
	obj := &models.Company{}
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
func (dao CompanyDAOPsql) GetAll() ([]models.Company, error) {
	query := `SELECT id, identification_type_id, identification_number, verification_digit, company, address, phone, departments_id, cities_id, web, email, activity, autorretenedor, person_type_id, regime_type_id, taxpayer_type_id, logo, created_at, updated_at
				FROM companies ORDER BY id`
	objs := make([]models.Company, 0)
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
		var o models.Company
		err = rows.Scan(&o.ID, &o.IdentificationType.ID, &o.IdentificationNumber, &o.VerificationDigit, &o.Company, &o.Address, &o.Phone, &o.Department.ID, &o.City.ID, &o.Web, &o.Email, &o.Activity, &o.AutoRretenedor, &o.PersonType.ID, &o.RegimeType.ID, &o.TaxpayerType.ID, &o.Logo, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao CompanyDAOPsql) rowToObject(row *sql.Row, o *models.Company) error {
	return row.Scan(&o.ID, &o.IdentificationType.ID, &o.IdentificationNumber, &o.VerificationDigit, &o.Company, &o.Address, &o.Phone, &o.Department.ID, &o.City.ID, &o.Web, &o.Email, &o.Activity, &o.AutoRretenedor, &o.PersonType.ID, &o.RegimeType.ID, &o.TaxpayerType.ID, &o.Logo, &o.CreatedAt, &o.UpdatedAt)
}
