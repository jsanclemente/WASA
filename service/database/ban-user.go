package database

import "database/sql"

// "banner" bans user "banned". If everything is okey, returns the number of users banned for the user
func (db *appdbimpl) BanUser(banner uint64, banned uint64) (uint64, error){

	var idUser uint64
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		banner).Scan(&idUser); err != nil { //Check if "banner" exists
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}
	// "banner" exists. We have to check "banned"
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		banned).Scan(&idUser); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}

	// At this point, both "banner" and "banned" exists. Insert to table Bans
	_, err := db.c.Exec(`INSERT INTO Bans (banner_id,banned_id) VALUES (?, ?)`,
		banner, banned)
	if err != nil {
		return 0, err
	}

	// Comprobar si se siguen, para dejarse de seguir
	// User banner stops following banned
	db.UnfollowUser(banner,banned)
	//Banned stops following 
	db.UnfollowUser(banner,banned)


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

	return nBans, nil
}