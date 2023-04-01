package database

// "unfollower" unfollows the user "unfollowed". If everything is okey returns the number of users followed by "unfollower",
// otherwise, returns 0.
func (db *appdbimpl) UnfollowUser(unfollower uint64, unfollowed uint64) (uint64, error) {

	//Check if unfollower exists
	if !db.UserExists(unfollower) {
		return 0, UserSubjectNotExists
	}
	//Check if unfollowed exists
	if !db.UserExists(unfollowed) {
		return 0, UserPredicateNotExists
	}

	// At this point, both "unfollower" and "unfollowed" exists. Delete on table Follows
	if !db.IsFollowing(unfollower, unfollowed) {
		return 0, ErrUser1alreadyFollows2
	}
	_, err := db.c.Exec(`DELETE FROM Follows WHERE follower_id=? AND followed_id=?`,
		unfollower, unfollowed)
	if err != nil {
		return 0, err
	}

	// Buscar el usuario que deja de seguir por el id, decrementar sus seguidos
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
