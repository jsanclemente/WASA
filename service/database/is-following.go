package database

import "database/sql"

func (db *appdbimpl) IsFollowing(follower uint64, followed uint64) bool {

	var idUser uint64
	if err := db.c.QueryRow("SELECT follower_id FROM Follows where follower_id = ? AND followed_id = ?",
		follower, followed).Scan(&idUser); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
