package postgresql

import (
	"github.com/alexyslozada/accounting-go/models"
	"bytes"
)

// PermissionDAOPsql consulta si el perfil tiene o no permisos
type PermissionDAOPsql struct {}

func (dao PermissionDAOPsql) IsPermitted(profile models.Profile, path string, method string) (bool, error) {
	var action, query string
	var result bool

	bufferQuery := bytes.NewBufferString("SELECT ")
	switch method {
	case "GET":
		action = "get"
	case "POST":
		action = "post"
	case "PUT":
		action = "put"
	case "DELETE":
		action = "del"
	}
	bufferQuery.WriteString(action)
	bufferQuery.WriteString(" FROM path_profile WHERE profile_id = $1 AND path_id = (SELECT id FROM paths WHERE path = $2)")

	query = bufferQuery.String()

	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(profile.ID, path).Scan(&result)
	return result, err
}
