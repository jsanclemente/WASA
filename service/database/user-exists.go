package database

import "database/sql"

// If the user exists the fucntion returns true, otherwise, returns false.
func (db *appdbimpl) UserExists(user uint64) bool {
	var idUser uint64
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		user).Scan(&idUser); err != nil { // Check if "user" exists
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
