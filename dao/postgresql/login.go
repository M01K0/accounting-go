package postgresql

import "github.com/alexyslozada/accounting-go/models"

type LoginDAOPsql struct {}

func (l LoginDAOPsql) Login(u *models.User) error {
	query := `SELECT users.id, identification, username, profile_id, profile
				FROM users INNER JOIN profiles ON users.profile_id = profiles.id
				WHERE email = $1 AND passwd = md5($2) AND users.active = true AND profiles.active = true`
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.Email, u.Passwd)
	return row.Scan(&u.ID, &u.Identification, &u.Username, &u.Profile.ID, &u.Profile.Profile)
}
