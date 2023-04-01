package database

import "database/sql"

// If "photoId" returns true, otherwise returns false
func (db *appdbimpl) PhotoExists(photoId uint64) bool {
	var id uint64
	if err := db.c.QueryRow("SELECT id FROM Photos where id = ?",
		photoId).Scan(&id); err != nil { //Check if "photoId" exists
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
