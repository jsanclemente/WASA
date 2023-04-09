package database

// Returns true if and only if username is not already registered on de DB
func (db *appdbimpl) IsValid(username string) bool {
	var count int
	if err := db.c.QueryRow("SELECT count(*) FROM Users where username = ?",
		username).Scan(&count); err != nil {
		if count == 0 {
			return true
		}
	}
	return false
}
