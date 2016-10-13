package postgresql

import (
	"fmt"
	"strings"
)

// PermissionDAOPsql consulta si el perfil tiene o no permisos
type PermissionDAOPsql struct{}

func (dao PermissionDAOPsql) GetScopes(id int16, method string) ([]string, error) {
	query := fmt.Sprintf(`
		SELECT paths.path
		FROM path_profile INNER JOIN paths ON path_profile.id = paths.id
		WHERE path_profile.profile_id = $1 AND %s = true
		ORDER BY paths.path
	`, strings.ToLower(method))

	scopes := make([]string, 0)
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
		var path string
		err = rows.Scan(&path)
		if err != nil {
			return scopes, err
		}

		scopes = append(scopes, path)
	}
	return scopes, nil
}
