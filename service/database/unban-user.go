package database

// "unbanner" bans user "unbanned". If everything is okey, returns the number of users unbanned for the user
func (db *appdbimpl) UnbanUser(unbanner uint64, unbanned uint64) (uint64, error) {

	// Check if unbanner exists
	if !db.UserExists(unbanner) {
		return 0, UserSubjectNotExists
	}
	//Check if unbanned exists
	if !db.UserExists(unbanned) {
		return 0, UserPredicateNotExists
	}

	// At this point, both "unbanner" and "unbanned" exists. Delete on table Bans
	_, err := db.c.Exec(`DELETE FROM Bans WHERE banner_id=? AND banned_id=?`,
		unbanner, unbanned)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario banneador por el id, decrementar sus baneados
	var nBans uint64
	if err := db.c.QueryRow("SELECT nbans FROM Users where id = ?",
		unbanner).Scan(&nBans); err != nil {
		return 0, err
	}
	nBans = nBans - 1
	_, err = db.c.Exec(`UPDATE Users SET nbans=? WHERE id=?`,
		nBans, unbanner)
	if err != nil {
		return 0, err
	}

	return nBans, nil
}
