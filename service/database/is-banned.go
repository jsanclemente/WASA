package database

import "database/sql"

// Returns true if and only if a is banned by b
func (db *appdbimpl) IsBanned(a uint64, b uint64) bool {
	var bannerId uint64
	if err := db.c.QueryRow("SELECT banner_id FROM Bans where banner_id = ? AND banned_id = ?",
		b, a).Scan(&bannerId); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
