package database

import "database/sql"

// "unbanner" bans user "unbanned". If everything is okey, returns the number of users unbanned for the user
func (db *appdbimpl) UnbanUser(unbanner uint64, unbanned uint64) (uint64, error) {
	var idUser uint64
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		unbanner).Scan(&idUser); err != nil { //Check if "unbanner" exists
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}
	// "unbanner" exists. We have to check "unbanned"
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		unbanned).Scan(&idUser); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
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
