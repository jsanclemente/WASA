package database

// Function to search users that starts like the parameter "username"
func (db *appdbimpl) SearchUser(username string, query string) ([]User, error) {

	// Obtener las publicaciones del usuario
	rows, err := db.c.Query("SELECT * FROM Users WHERE username LIKE ? AND username != ? ", query+"%", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		var nBans uint64
		err := rows.Scan(&user.ID, &user.Username, &user.Nfollowers, &user.Nfollowing, &user.Nposts, &nBans)
		if err != nil {
			return nil, err
		}
		user.Followers = nil
		user.Following = nil
		user.Posts = nil

		users = append(users, user)
	}

	// Check if there were errors during rows.Next()
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
