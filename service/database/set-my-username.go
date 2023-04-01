package database

import "database/sql"

// The username associated to user "userId" changes to "username". Returns the old username.
// If an error occurs, returns the empty string

func (db *appdbimpl) SetMyUserName(userId uint64, username string) (string, error) {
	// 1.Comprobar que existe el usuario
	// 1.1 Si existe modificar el campo y devolver el antiguo

	// 1.
	var oldUsername string
	if err := db.c.QueryRow("SELECT username FROM Users where id = ?",
		userId).Scan(&oldUsername); err != nil {
		if err == sql.ErrNoRows {
			return "", UserSubjectNotExists
		}
	}

	_, err := db.c.Exec(`UPDATE Users SET username=? WHERE id=?`,
		username, userId)
	if err != nil {
		return "", err
	}

	return oldUsername, nil
}
