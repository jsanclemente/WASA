package database

// Given a id for one photo, returns the binary imagefor that photo
func (db *appdbimpl) GetPhoto(photoId uint64) (Photo, error) {

	// 1. Check if the photo exists
	if !db.PhotoExists(photoId) {
		return Photo{}, ErrPhotoNotExits
	}

	// 2. Obtain the information of the photo
	var photo Photo

	if err := db.c.QueryRow("SELECT p.id, p.nLikes, p.nComments, p.imageData, Posts.date FROM Photos p INNER JOIN Posts ON p.id=Posts.photo_id WHERE Posts.photo_id=?",
		photoId).Scan(&photo.ID, &photo.Nlikes, &photo.Ncomments, &photo.Image, &photo.Date); err != nil {
		return Photo{}, err
	}

	rows, err := db.c.Query("SELECT comment_id, comment, username FROM Comments c JOIN Users u ON c.user_id = u.id WHERE photo_id = ? ORDER BY comment_id ASC", photoId)
	if err != nil {
		return Photo{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.IdComment, &comment.Comment, &comment.Username); err != nil {
			return Photo{}, err
		}
		photo.Comments = append(photo.Comments, comment.IdComment)
	}
	if err := rows.Err(); err != nil {
		return Photo{}, err
	}

	return photo, err
}
