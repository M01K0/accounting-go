package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
)

// PermissionDAOPsql consulta si el perfil tiene o no permisos
type PermissionDAOPsql struct {}

func (dao PermissionDAOPsql) GetScopes(id int16) ([]models.Scope, error) {
	query := `
		SELECT paths.path, post, put, del, get
		FROM path_profile INNER JOIN paths ON path_profile.id = paths.id
		WHERE path_profile.profile_id = $1 AND (post = true OR put = true OR del = true OR get = true)
		ORDER BY paths.path
	`

	scopes := make([]models.Scope, 0)
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return scopes, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var path string
	var post, put, del, get bool
	for rows.Next() {
		var scope models.Scope
		err = rows.Scan(&path, &post, &put, &del, &get)
		if err != nil {
			return scopes, err
		}

		scope.Path = path
		if post {
			scope.Methods = append(scope.Methods, "POST")
		}
		if put {
			scope.Methods = append(scope.Methods, "PUT")
		}
		if del {
			scope.Methods = append(scope.Methods, "DELETE")
		}
		if get {
			scope.Methods = append(scope.Methods, "GET")
		}
		scopes = append(scopes, scope)
	}
	return scopes, nil
}
