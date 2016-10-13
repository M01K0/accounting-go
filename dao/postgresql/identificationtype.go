package postgresql

import (
	"database/sql"
	"errors"
	"github.com/alexyslozada/accounting-go/models"
)

type IdentificationTypeDAOPsql struct{}

// Insert insertar registro en la BD
func (dao IdentificationTypeDAOPsql) Insert(obj *models.IdentificationType) error {
	query := "INSERT INTO identification_type (initials, identification_name, dian_code) VALUES (upper($1), upper($2), upper($3)) RETURNING id, initials, identification_name, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.Initials, obj.IdentificationName, obj.DianCode)
	return dao.rowToObject(row, obj)
}

// Update actualizar registro en la bd
func (dao IdentificationTypeDAOPsql) Update(obj *models.IdentificationType) error {
	query := "UPDATE identification_type SET initials = upper($2), identification_name = upper($3), dian_code = upper($4), updated_at = now() WHERE id = $1 RETURNING id, initials, identification_name, dian_code, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow(obj.ID, obj.Initials, obj.IdentificationName, obj.DianCode)
	return dao.rowToObject(row, obj)
}

// Delete borrar registro de la bd
func (dao IdentificationTypeDAOPsql) Delete(obj *models.IdentificationType) error {
	query := "DELETE FROM identification_type WHERE id = $1"
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
	obj = new(models.IdentificationType)
	return nil
}

// GetByID consultar registro por id
func (dao IdentificationTypeDAOPsql) GetByID(id int) (*models.IdentificationType, error) {
	query := "SELECT id, initials, identification_name, dian_code, created_at, updated_at FROM identification_type WHERE id = $1"
	obj := &models.IdentificationType{}
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
func (dao IdentificationTypeDAOPsql) GetAll() ([]models.IdentificationType, error) {
	query := "SELECT id, initials, identification_name, dian_code, created_at, updated_at FROM identification_type ORDER BY id"
	objs := make([]models.IdentificationType, 0)
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
		var o models.IdentificationType
		err = rows.Scan(&o.ID, &o.Initials, &o.IdentificationName, &o.DianCode, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return objs, err
		}
		objs = append(objs, o)
	}
	return objs, nil
}

func (dao IdentificationTypeDAOPsql) rowToObject(row *sql.Row, o *models.IdentificationType) error {
	return row.Scan(&o.ID, &o.Initials, &o.IdentificationName, &o.DianCode, &o.CreatedAt, &o.UpdatedAt)
}
