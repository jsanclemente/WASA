package database

// Given a id for one user, returns the users followed by "userId"
func (db *appdbimpl) GetFollowing(userId uint64) ([]User, error) {

	// 1. Check if the user exists
	if !db.UserExists(userId) {
		return nil, ErrUserSubjectNotExists
	}

	// 2. Obtain the following for "userId"
	var users []User
	rows, err := db.c.Query("SELECT followed_id, username FROM Users u JOIN Follows f ON u.id = f.followed_id WHERE f.follower_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
