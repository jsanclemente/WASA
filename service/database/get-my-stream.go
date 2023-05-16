package database

import (
	"time"
)

// Given a username for one user, returns the stream for that user.
func (db *appdbimpl) GetMyStream(userId uint64) ([]Photo, error) {

	// Por cada usuario que sigue userId, mirar en la tabla posts que fotos tiene subidas y despues mirar en la tabla fotos para sacar los likes y comments
	// obtener los usuarios a los que sigue el idUser

	//Check if the user exists
	if !db.UserExists(userId) {
		return nil, ErrUserSubjectNotExists
	}

	// ----------------------------------------------------------------------------------

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
		rows, err := db.c.Query("SELECT p.id, p.nLikes, p.nComments, p.imageData, Posts.date FROM Photos p INNER JOIN Posts ON p.id=Posts.photo_id WHERE Posts.user_id=?", followedUser)
		if err != nil {
			return nil, err
		}
		defer func() { _ = rows.Close() }()

		var username string

		// Extract the username
		if err := db.c.QueryRow("SELECT username FROM Users where id = ?",
			followedUser).Scan(&username); err != nil {
			return nil, err
		}

		// Para cada publicacion, obtener sus campos
		var post Photo
		post.Username = username
		for rows.Next() {
			post.Comments = nil
			post.Likes = nil

			if err := rows.Scan(&post.ID, &post.Nlikes, &post.Ncomments, &post.Image, &post.Date); err != nil {
				return nil, err
			}

			// Conver to type "time"
			datetype, err := time.Parse("2006-01-02T15:04:05Z", post.Date)
			if err != nil {
				return nil, err
			}
			post.Date = datetype.Format("02-01-2006") //Sacar fecha
			post.Time = datetype.Format("15:04:05")   //Sacar la hora
			if err != nil {
				return nil, err
			}

			//	Obtener los id's de los usuarios que han dado like para cada publicación
			rows, err := db.c.Query("SELECT user_id FROM Likes WHERE photo_id = ?", post.ID)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			for rows.Next() {
				var user uint64
				if err := rows.Scan(&user); err != nil {
					return nil, err
				}
				post.Likes = append(post.Likes, user)
			}

			// Se obtienen los id's de los comentarios para cada publicación
			rows, err = db.c.Query("SELECT comment_id FROM Comments WHERE photo_id=?", post.ID)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			for rows.Next() {
				var commentId uint64
				if err := rows.Scan(&commentId); err != nil {
					return nil, err
				}
				post.Comments = append(post.Comments, commentId)
			}
			// -----------------------------------------------------------------------------------
			posts = append(posts, post)
		}
	}
	return posts, nil
}
