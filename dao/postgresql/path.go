package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"errors"
	"database/sql"
)

// PathDAOPsql estructura dao de Path
type PathDAOPsql struct {}

// Insert insertar
func (dao PathDAOPsql) Insert(path *models.Path) error {
	query := "INSERT INTO paths (path, path_name, description) VALUES (lower($1), upper($2), $3) RETURNING id, code, path_name, description, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRow(path.Path, path.PathName, path.Description)
	return dao.rowToPath(row, path)
}

// Update actualizar
func (dao PathDAOPsql) Update(path *models.Path) error {
	query := "UPDATE paths SET path = lower($1), path_name = upper($2), description = upper($3), updated_at = now() WHERE id = $4 RETURNING id, path, path_name, description, created_at, updated_at"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(path.Path, path.PathName, path.Description, path.ID)
	return dao.rowToPath(row, path)

}

// Delete borrar
func (dao PathDAOPsql) Delete(path *models.Path) error {
	query := "DELETE FROM paths WHERE id = $1"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(path.ID)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return errors.New("No se eliminó ningún registro")
	}
	path = new(models.Path)
	return nil
}

// GetByID Consulta por id
func (dao PathDAOPsql) GetByID(id int) (*models.Path, error) {
	query := "SELECT id, path, path_name, description, created_at, updated_at FROM paths WHERE id = $1"
	path := &models.Path{}
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = dao.rowToPath(row, path)
	return path, err
}

// GetAll Consulta todos
func (dao PathDAOPsql) GetAll() ([]models.Path, error) {
	query := "SELECT id, path, path_name, description, created_at, updated_at FROM paths ORDER BY path"
	paths := make([]models.Path, 0)
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
		var path models.Path
		err = rows.Scan(&path.ID, &path.Path, &path.PathName, &path.Description, &path.CreatedAt, &path.UpdatedAt)
		if err != nil {
			return paths, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}

// rowToPath mapea la consulta al objeto
func (dao PathDAOPsql) rowToPath(row *sql.Row, path *models.Path) error {
	return row.Scan(&path.ID, &path.Path, &path.PathName, &path.Description, &path.CreatedAt, &path.UpdatedAt)
}
