package database

import "database/sql"

// "follower" follows the user "followed". If everything is okey returns the number of users followed,
// otherwise, returns 0.
func (db *appdbimpl) FollowUser(follower uint64, followed uint64) (uint64, error) {

	// Comprobar que los dos usuarios existen.
	// Si existen:
	//	1. Añadir (follower,followed) a la tabla Follows
	//	2. Incrementar en 1, el numero de seguidos de follower
	//	3. Incrementar en 1, el numero de seguidores de followed

	var idUser uint64
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		follower).Scan(&idUser); err != nil { //Check if "follower" exists
		if err == sql.ErrNoRows {
			return 0, ErrUserSubjectNotExists
		}
		return 0, err
	}
	// "follower" exists. We have to check "followed"
	if err := db.c.QueryRow("SELECT id FROM Users where id = ?",
		followed).Scan(&idUser); err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrUserPredicateNotExists
		}
		return 0, err
	}

	if db.IsBanned(follower, followed) {
		return 0, ErrUserAIsBanned
	}
	if db.IsBanned(followed, follower) {
		return 0, ErrUserBIsBanned
	}
	// Check if "follower" already follows "followed"
	if db.IsFollowing(follower, followed) {
		return 0, ErrUser1alreadyFollows2
	}

	// At this point, both "follower" and "followed" exists. Insert to table Follows
	_, err := db.c.Exec(`INSERT INTO Follows (follower_id,followed_id) VALUES (?, ?)`,
		follower, followed)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario seguidor por el id, incrementar sus seguidos
	var nFollowing uint64
	if err := db.c.QueryRow("SELECT nfollowing FROM Users where id = ?",
		follower).Scan(&nFollowing); err != nil {
		return 0, err
	}
	nFollowing = nFollowing + 1
	_, err = db.c.Exec(`UPDATE Users SET nfollowing=? WHERE id=?`,
		nFollowing, follower)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario seguido por id, incrementar sus seguidores
	var nFollowers uint64
	if err := db.c.QueryRow("SELECT nfollowers FROM Users where id = ?",
		followed).Scan(&nFollowers); err != nil {
		return 0, err
	}
	nFollowers = nFollowers + 1
	_, err = db.c.Exec(`UPDATE Users SET nfollowers=? WHERE id=?`,
		nFollowers, followed)
	if err != nil {
		return 0, err
	}

	return nFollowing, nil
}
