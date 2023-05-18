package database

func (db *appdbimpl) GetUserProfile(userId uint64) (User, error) {
	// 1.Comprobar si existe el usuario
	//	1.1 Si existe, obtener todos los campos del usuario

	// 1.Check if the user exists
	if !db.UserExists(userId) {
		return User{}, ErrUserSubjectNotExists
	}

	var user User
	err := db.c.QueryRow("SELECT username, nfollowers, nfollowing, nposts FROM Users WHERE id = ?", userId).Scan(&user.Username, &user.Nfollowers, &user.Nfollowing, &user.Nposts)
	if err != nil {
		return User{}, err
	}

	user.ID = userId

	// Obtener las publicaciones del usuario
	rows, err := db.c.Query("SELECT photo_id FROM Posts WHERE user_id = ?", userId)
	if err != nil {
		return User{}, err
	}

	defer rows.Close()

	// Obtener la lista con todos los id's de las posts del usuario
	for rows.Next() {
		var postID int
		err = rows.Scan(&postID)
		if err != nil {
			return User{}, err
		}
		user.Posts = append(user.Posts, postID)
	}
	// ----------------------------------------------------------------------------

	// Obtener la lista con los id's de los usuarios a los que sigo
	rows, err = db.c.Query("SELECT followed_id FROM Follows WHERE follower_id = ?", userId)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var followingID int
		err = rows.Scan(&followingID)
		if err != nil {
			return User{}, err
		}
		user.Following = append(user.Following, followingID)
	}
	// ---------------------------------------------------------------------------------------

	// Obtener una lista de los usuarios que me siguen
	rows, err = db.c.Query("SELECT follower_id FROM Follows WHERE followed_id = ?", userId)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var followerID int
		err = rows.Scan(&followerID)
		if err != nil {
			return User{}, err
		}
		user.Followers = append(user.Followers, followerID)
	}

	// ---------------------------------------------------------------------------------------

	// Obtener una lista con los id's de usuario que tienen baneado a "userId"

	rows, err = db.c.Query("SELECT banner_id FROM Bans WHERE banned_id = ?", userId)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var bannerID int
		err = rows.Scan(&bannerID)
		if err != nil {
			return User{}, err
		}
		user.Banners = append(user.Banners, bannerID)
	}

	// ---------------------------------------------------------------------------------------
	return user, nil
}
