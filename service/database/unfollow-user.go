package database

import "database/sql"

// "unfollower" unfollows the user "unfollowed". If everything is okey returns the number of users followed,
// otherwise, returns 0.
func (db *appdbimpl) UnfollowUser(unfollower uint64, unfollowed uint64) (uint64, error) {

	var idUser uint64
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		unfollower).Scan(&idUser); err != nil { //Check if "unfollower" exists
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}
	// "unfollower" exists. We have to check "unfollowed"
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		unfollowed).Scan(&idUser); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}

	// At this point, both "unfollower" and "unfollowed" exists. Delete on table Follows
	_, err := db.c.Exec(`DELETE FROM Follows WHERE follower_id=? AND followed_id=?`,
		unfollower, unfollowed)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario seguidor por el id, decrementar sus seguidos
	var nFollowing uint64
	if err := db.c.QueryRow("SELECT nfollowing FROM Users where id = ?",
		unfollower).Scan(&nFollowing); err != nil {
		return 0, err
	}
	nFollowing = nFollowing - 1
	_, err = db.c.Exec(`UPDATE Users SET nfollowing=? WHERE id=?`,
		nFollowing, unfollower)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario seguido por id, incrementar sus seguidores
	var nFollowers uint64
	if err := db.c.QueryRow("SELECT nfollowers FROM Users where id = ?",
		unfollowed).Scan(&nFollowers); err != nil {
		return 0, err
	}
	nFollowers = nFollowers - 1
	_, err = db.c.Exec(`UPDATE Users SET nfollowers=? WHERE id=?`,
		nFollowers, unfollowed)
	if err != nil {
		return 0, err
	}

	return nFollowing, nil
}
