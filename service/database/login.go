package database

import "database/sql"

// Returns the id of the user, if the value returned is 0, something has gone wrong
func (db *appdbimpl) Login(username string) (uint64, error) {
	
	var idUser uint64

	if err := db.c.QueryRow("SELECT id FROM Users where username = ?",
		username).Scan(&idUser); err != nil {
		if err == sql.ErrNoRows {
			res, err := db.c.Exec(`INSERT INTO Users (id, username, nfollowers, nfollowing, nposts, nabans) VALUES (NULL, ?, ?, ?, ?, ?)`,
				username, 0, 0, 0, 0)
			if err != nil {
				return 0, err
			}
			lastInsertID, err := res.LastInsertId()
			if err != nil {
				return 0, err
			}

			idUser = uint64(lastInsertID)
			return idUser, nil
		}
		return 0, nil
	}
	return idUser, nil
}
