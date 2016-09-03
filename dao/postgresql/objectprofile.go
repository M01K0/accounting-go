package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"database/sql"
)

// ObjectProfileDAOPsql estructura dao de object_profile
type ObjectProfileDAOPsql struct {}

// Update actualiza un registro en la BD
func (dao ObjectProfileDAOPsql) Update(o *models.ObjectProfile) error {
	query := "UPDATE object_profile SET creates = $1, modify = $2, erase = $3, query = $4 updated_at = now() WHERE id = $5 RETURNING id, profile_id, object_id, creates, modify, erase, query, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(o.Creates, o.Modify, o.Erase, o.Query, o.ID)
	return dao.rowToObjectProfile(row, o)
}

// GetByID obtiene un registro de la BD
func (dao ObjectProfileDAOPsql) GetByID(id int) (*models.ObjectProfile, error) {
	query := "SELECT id, profile_id, object_id, creates, modify, erase, query, created_at, updated_at FROM object_profile WHERE id = $1"
	o := &models.ObjectProfile{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = dao.rowToObjectProfile(row, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// GetByProfileID obtiene el registro por el id del perfil
func (dao ObjectProfileDAOPsql) GetByProfileID(id int16) ([]models.ObjectProfile, error) {
	query := `SELECT op.id, op.profile_id, op.object_id, o.code, o.object_name, o.description, op.creates,
					op.modify, op.erase, op.query, op.created_at, op.updated_at
					FROM object_profile AS op INNER JOIN objects AS o ON op.object_id = o.id
					WHERE profile_id = $1 ORDER BY o.code`

	ops := make([]models.ObjectProfile, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.ObjectProfile
		err = rows.Scan(&o.ID, &o.ProfileID, &o.Object.ID, &o.Object.Code, &o.Object.ObjectName, &o.Object.Description, &o.Creates, &o.Modify, &o.Erase, &o.Query, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return ops, err
		}
		ops = append(ops, o)
	}
	return ops, nil
}

// rowToObjectProfile mapea la consulta al objeto
func (dao ObjectProfileDAOPsql) rowToObjectProfile(row *sql.Row, o *models.ObjectProfile) error {
	return row.Scan(&o.ID, &o.ProfileID, &o.Object.ID, &o.Creates, &o.Modify, &o.Erase, &o.Query, &o.CreatedAt, &o.UpdatedAt)
}
