package database

import (
	"database/sql"
)

// Returns true if and only if "userId" has posted the post "photoId"
func (db *appdbimpl) OwnsPhoto(userId uint64, photoId uint64) bool {
	var id uint64
	if err := db.c.QueryRow("SELECT user_id FROM Posts where user_id = ? AND photo_id = ?",
		userId, photoId).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
