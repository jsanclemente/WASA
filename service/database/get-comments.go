package database

// Given a id for one photo, returns the comments for that photo
func (db *appdbimpl) GetComments(photoId uint64) ([]Comment, error) {

	// 1. Check if the photo exists
	if !db.PhotoExists(photoId) {
		return nil, ErrPhotoNotExits
	}

	// 2. Obtain the list
	rows, err := db.c.Query("SELECT comment_id, comment, username FROM Comments c JOIN Users u ON c.user_id = u.id WHERE photo_id = ? ORDER BY comment_id ASC", photoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.IdComment, &comment.Comment, &comment.Username); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
