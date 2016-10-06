package postgresql

// PermissionDAOPsql consulta si el perfil tiene o no permisos
type PermissionDAOPsql struct {}

func (dao PermissionDAOPsql) GetScopes(id int16) (map[string][]string, error) {
	query := `
		SELECT paths.path, post, put, del, get
		FROM path_profile INNER JOIN paths ON path_profile.id = paths.id
		WHERE path_profile.profile_id = $1 AND (post = true OR put = true OR del = true OR get = true)
		ORDER BY paths.path
	`

	scopes := make(map[string][]string)
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

	var path string
	var post, put, del, get bool
	for rows.Next() {
		methods := make([]string, 0)
		err = rows.Scan(&path, &post, &put, &del, &get)
		if err != nil {
			return scopes, err
		}

		if post {
			methods = append(methods, "POST")
		}
		if put {
			methods = append(methods, "PUT")
		}
		if del {
			methods = append(methods, "DELETE")
		}
		if get {
			methods = append(methods, "GET")
		}
		scopes[path] = methods
	}
	return scopes, nil
}
