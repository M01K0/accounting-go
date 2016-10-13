package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type ThirdPartyDAOPsql struct{}

// Insert insertar registro en la BD
func (dao ThirdPartyDAOPsql) Insert(obj *models.ThirdParty) error {
	query := `INSERT INTO third_parties (identification_type_id, identification_number, verification_digit, person_type_id, regime_type_id, taxpayer_type_id, business_name, last_name, second_last_name, first_name, middle_name, address, phone, email, department_id, city_id)
				VALUES ($1, $2, $3, $4, $5, $6, upper($7), upper($8), upper($9), upper($10), upper($11), upper($12), upper($13), lower($14), $15, $16)
				RETURNING id, identification_type_id, identification_number, verification_digit, person_type_id, regime_type_id, taxpayer_type_id, business_name, last_name, second_last_name, first_name, middle_name, address, phone, email, department_id, city_id, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.IdentificationType.ID, obj.IdentificationNumber, obj.VerificationDigit, obj.PersonType.ID, obj.RegimeType.ID, obj.TaxpayerType.ID, obj.BusinessName, obj.LastName, obj.SecondLastName, obj.FirstName, obj.MiddleName, obj.Address, obj.Phone, obj.Email, obj.Department.ID, obj.City.ID)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao ThirdPartyDAOPsql) Update(obj *models.ThirdParty) error {
	query := `UPDATE third_parties
				SET identification_type_id = $2, identification_number = $3, verification_digit = $4, person_type_id = $5, regime_type_id = $6, taxpayer_type_id = $7, business_name = upper($8), last_name = upper($9), second_last_name = upper($10), first_name = upper($11), middle_name = upper($12), address = upper($13), phone = upper($14), email = lower($15), department_id = $16, city_id = $17, updated_at = now()
				WHERE id = $1
				RETURNING id, identification_type_id, identification_number, verification_digit, person_type_id, regime_type_id, taxpayer_type_id, business_name, last_name, second_last_name, first_name, middle_name, address, phone, email, department_id, city_id, created_at, updated_at`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.IdentificationType.ID, obj.IdentificationNumber, obj.VerificationDigit, obj.PersonType.ID, obj.RegimeType.ID, obj.TaxpayerType.ID, obj.BusinessName, obj.LastName, obj.SecondLastName, obj.FirstName, obj.MiddleName, obj.Address, obj.Phone, obj.Email, obj.Department.ID, obj.City.ID)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao ThirdPartyDAOPsql) Delete(obj *models.ThirdParty) error {
	query := "DELETE FROM third_parties WHERE id = $1"
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
	obj = new(models.ThirdParty)
	return nil
}

// GetByID consultar registro por id
func (dao ThirdPartyDAOPsql) GetByID(id int) (*models.ThirdParty, error) {
	query := `SELECT
				third_parties.id,
				identification_type.id AS identification_type_id,
				identification_type.initials, identification_type.identification_name,
				identification_type.dian_code,
				third_parties.identification_number, third_parties.verification_digit,
				person_type.id AS person_type_id, person_type.person,
				regime_type.id AS regime_type_id, regime_type.regime,
				taxpayer_type.id AS taxpayer_type_id, taxpayer_type.taxpayer,
				third_parties.business_name, third_parties.last_name,
				third_parties.second_last_name, third_parties.first_name,
				third_parties.middle_name, third_parties.address, third_parties.phone,
				third_parties.email,
				departments.id AS department_id, departments.code, departments.department,
				cities.id AS city_id, cities.code, cities.city,
				third_parties.created_at, third_parties.updated_at
			FROM
				third_parties INNER JOIN
				cities ON third_parties.city_id = cities.id INNER JOIN
				departments ON third_parties.department_id = departments.id INNER JOIN
				identification_type ON third_parties.identification_type_id = identification_type.id INNER JOIN
				person_type ON third_parties.person_type_id = person_type.id INNER JOIN
				regime_type ON third_parties.regime_type_id = regime_type.id INNER JOIN
				taxpayer_type ON third_parties.taxpayer_type_id = taxpayer_type.id
			WHERE id = $1`
	obj := &models.ThirdParty{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&obj.ID, &obj.IdentificationType.ID, &obj.IdentificationType.Initials, &obj.IdentificationType.IdentificationName, &obj.IdentificationType.DianCode, &obj.IdentificationNumber, &obj.VerificationDigit, &obj.PersonType.ID, &obj.PersonType.Person, &obj.RegimeType.ID, &obj.RegimeType.Regime, &obj.TaxpayerType.ID, &obj.TaxpayerType.Taxpayer, &obj.BusinessName, &obj.LastName, &obj.SecondLastName, &obj.FirstName, &obj.MiddleName, &obj.Address, &obj.Phone, &obj.Email, &obj.Department.ID, &obj.Department.Code, &obj.Department.Department, &obj.City.ID, &obj.City.Code, &obj.City.City, &obj.CreatedAt, &obj.UpdatedAt)
	return obj, err
}

// GetAll Consulta todos los registros de la bd
func (dao ThirdPartyDAOPsql) GetAll() ([]models.ThirdParty, error) {
	query := `SELECT
				third_parties.id,
				identification_type.id AS identification_type_id,
				identification_type.initials, identification_type.identification_name,
				identification_type.dian_code,
				third_parties.identification_number, third_parties.verification_digit,
				person_type.id AS person_type_id, person_type.person,
				regime_type.id AS regime_type_id, regime_type.regime,
				taxpayer_type.id AS taxpayer_type_id, taxpayer_type.taxpayer,
				third_parties.business_name, third_parties.last_name,
				third_parties.second_last_name, third_parties.first_name,
				third_parties.middle_name, third_parties.address, third_parties.phone,
				third_parties.email,
				departments.id AS department_id, departments.code, departments.department,
				cities.id AS city_id, cities.code, cities.city,
				third_parties.created_at, third_parties.updated_at
			FROM
				third_parties INNER JOIN
				cities ON third_parties.city_id = cities.id INNER JOIN
				departments ON third_parties.department_id = departments.id INNER JOIN
				identification_type ON third_parties.identification_type_id = identification_type.id INNER JOIN
				person_type ON third_parties.person_type_id = person_type.id INNER JOIN
				regime_type ON third_parties.regime_type_id = regime_type.id INNER JOIN
				taxpayer_type ON third_parties.taxpayer_type_id = taxpayer_type.id
			ORDER BY id`
	objs := make([]models.ThirdParty, 0)
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
		var obj models.ThirdParty
		err = rows.Scan(&obj.ID, &obj.IdentificationType.ID, &obj.IdentificationType.Initials, &obj.IdentificationType.IdentificationName, &obj.IdentificationType.DianCode, &obj.IdentificationNumber, &obj.VerificationDigit, &obj.PersonType.ID, &obj.PersonType.Person, &obj.RegimeType.ID, &obj.RegimeType.Regime, &obj.TaxpayerType.ID, &obj.TaxpayerType.Taxpayer, &obj.BusinessName, &obj.LastName, &obj.SecondLastName, &obj.FirstName, &obj.MiddleName, &obj.Address, &obj.Phone, &obj.Email, &obj.Department.ID, &obj.Department.Code, &obj.Department.Department, &obj.City.ID, &obj.City.Code, &obj.City.City, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

func (dao ThirdPartyDAOPsql) rowToObject(row *sql.Row, o *models.ThirdParty) error {
	return row.Scan(&o.ID, &o.IdentificationType.ID, &o.IdentificationNumber, &o.VerificationDigit, &o.PersonType.ID, &o.RegimeType.ID, &o.TaxpayerType.ID, &o.BusinessName, &o.LastName, &o.SecondLastName, &o.FirstName, &o.MiddleName, &o.Address, &o.Phone, &o.Email, &o.Department.ID, &o.City.ID, &o.CreatedAt, &o.UpdatedAt)
}
