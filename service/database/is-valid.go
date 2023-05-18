package database

// Returns true if and only if username is not already registered on de DB
func (db *appdbimpl) IsValid(username string) bool {
	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Users where username = ?", username).Scan(&count)
	if err == nil {
		if count == 0 {
			return true
		}
	} else {
		return false
	}

	return false
}
