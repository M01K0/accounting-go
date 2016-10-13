package postgresql

import (
	"database/sql"
	"github.com/alexyslozada/accounting-go/models"
)

// PathProfileDAOPsql estructura dao de path_profile
type PathProfileDAOPsql struct{}

// Update actualiza un registro en la BD
func (dao PathProfileDAOPsql) Update(o *models.PathProfile) error {
	query := "UPDATE path_profile SET post = $1, put = $2, del = $3, get = $4 updated_at = now() WHERE id = $5 RETURNING id, profile_id, path_id, post, put, del, get, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(o.Post, o.Put, o.Del, o.Get, o.ID)
	return dao.rowToPathProfile(row, o)
}

// GetByID obtiene un registro de la BD
func (dao PathProfileDAOPsql) GetByID(id int) (*models.PathProfile, error) {
	query := "SELECT id, profile_id, path_id, post, put, del, get, created_at, updated_at FROM path_profile WHERE id = $1"
	o := &models.PathProfile{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = dao.rowToPathProfile(row, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// GetByProfileID obtiene el registro por el id del perfil
func (dao PathProfileDAOPsql) GetByProfileID(id int16) ([]models.PathProfile, error) {
	query := `SELECT pp.id, pp.profile_id, pp.path_id, pa.path, pa.path_name,
					pa.description, pp.post, pp.put, pp.del, pp.get,
					pp.created_at, pp.updated_at
				FROM path_profile AS pp INNER JOIN paths AS pa ON pp.path_id = pa.id
				WHERE profile_id = $1 ORDER BY pa.path`

	ops := make([]models.PathProfile, 0)
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
		var o models.PathProfile
		err = rows.Scan(&o.ID, &o.ProfileID, &o.Path.ID, &o.Path.Path, &o.Path.PathName, &o.Path.Description, &o.Post, &o.Put, &o.Del, &o.Get, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return ops, err
		}
		ops = append(ops, o)
	}
	return ops, nil
}

// rowToPathProfile mapea la consulta al objeto
func (dao PathProfileDAOPsql) rowToPathProfile(row *sql.Row, o *models.PathProfile) error {
	return row.Scan(&o.ID, &o.ProfileID, &o.Path.ID, &o.Path.Path, &o.Path.PathName, &o.Path.Description, &o.Post, &o.Put, &o.Del, &o.Get, &o.CreatedAt, &o.UpdatedAt)
}
