package database

import "fmt"

// Function to search users that starts like the parameter "username"
func (db *appdbimpl) SearchUser(username string, query string, id uint64) ([]User, error) {

	// Obtener las publicaciones del usuario
	rows, err := db.c.Query("SELECT * FROM Users WHERE username LIKE ? AND username != ? AND id NOT IN (SELECT banner_id FROM Bans WHERE banned_id = ?)", query+"%", username, id, id)

	if err != nil {
		fmt.Println("error")
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
