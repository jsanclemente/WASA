package database

// Given a username for one user, returns the stream for that user.
func (db *appdbimpl) GetMyStream(userId uint64) ([]Photo, error) {
	// Por cada usuario que sigue userId, mirar en la tabla posts que fotos tiene subidas y despues mirar en la tabla fotos para sacar los likes y comments
	// obtener los usuarios a los que sigue el idUser

	//Check if the user exists
	if !db.UserExists(userId) {
		return nil, UserSubjectNotExists
	}

	// 1.Obtener la lista de usuarios a los que sigue "userId"
	rows, err := db.c.Query("SELECT followed_id FROM Follows WHERE follower_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	//Se almacenan todos los usuarios a los que sigue el "userId" en una lista de id's
	var followedUsers []int
	for rows.Next() {
		var followedUser int
		if err := rows.Scan(&followedUser); err != nil {
			return nil, err
		}
		followedUsers = append(followedUsers, followedUser)
	}

	if rr := rows.Err(); rr != nil {
		return nil, nil
	}
	//	---------------------------------------------------------------------------------------------

	// 2. Obtener las publicaciones de cada usuario seguido
	var posts []Photo
	for _, followedUser := range followedUsers {
		rows, err := db.c.Query("SELECT p.id, p.nLikes, p.nComments, p.url, Posts.date FROM Photos p INNER JOIN Posts ON p.id=Posts.photo_id WHERE Posts.user_id=?", followedUser)
		if err != nil {
			return nil, err
		}

		defer func() { _ = rows.Close() }()

		// Para cada publicacion, obtener sus campos
		var post Photo
		for rows.Next() {
			if err := rows.Scan(&post.ID, &post.nLikes, &post.nComments, &post.url, &post.date); err != nil {
				return nil, err
			}

			// Se obtienen los id's de los comentarios para cada publicación
			rows, err := db.c.Query("SELECT comment_id FROM Comments WHERE photo_id=?", post.ID)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			for rows.Next() {
				var commentId uint64
				if err := rows.Scan(&commentId); err != nil {
					return nil, nil
				}
				post.comments = append(post.comments, commentId)
			}
			// -----------------------------------------------------------------------------------
			posts = append(posts, post)
		}
	}
	return posts, nil
}
