package database

import "database/sql"

// Returns true if and only if userId has already liked photoId
func (db *appdbimpl) AlreadyLiked(userId uint64, photoId uint64) bool {
	var n uint64
	if err := db.c.QueryRow("SELECT user_id FROM Likes where user_id = ? AND photo_id = ?",
		userId, photoId).Scan(&n); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
