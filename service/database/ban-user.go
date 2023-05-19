package database

// "banner" bans user "banned". If everything is okey, returns the number of users banned for the user
func (db *appdbimpl) BanUser(banner uint64, banned uint64) (uint64, error) {

	//	Check if banner and banned exists
	if !db.UserExists(banner) {
		return 0, ErrUserSubjectNotExists
	}
	if !db.UserExists(banned) {
		return 0, ErrUserPredicateNotExists
	}

	//	Chequear if "banned" is already banned by "banner"
	if db.IsBanned(banned, banner) {
		return 0, ErrAlreadyBanned
	}

	// At this point, both "banner" and "banned" exists. Insert to table Bans
	_, err := db.c.Exec(`INSERT INTO Bans (banner_id,banned_id) VALUES (?, ?)`,
		banner, banned)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario banneador por el id, incrementar sus baneados
	var nBans uint64
	if err := db.c.QueryRow("SELECT nbans FROM Users where id = ?",
		banner).Scan(&nBans); err != nil {
		return 0, err
	}
	nBans = nBans + 1
	_, err = db.c.Exec(`UPDATE Users SET nbans=? WHERE id=?`,
		nBans, banner)
	if err != nil {
		return 0, err
	}

	//	Si banner seguia a banned, dejarlo de seguir.
	if db.IsFollowing(banner, banned) {
		_, err := db.UnfollowUser(banner, banned)
		if err != nil {
			return 0, err
		}
	}
	//	Si banned seguia a banner, que banned deje de seguirlo
	if db.IsFollowing(banned, banner) {
		_, err := db.UnfollowUser(banned, banner)
		if err != nil {
			return 0, err
		}
	}

	return nBans, nil
}
