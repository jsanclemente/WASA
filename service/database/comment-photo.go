package database

import (
	"database/sql"
)

// The user identified by "userId" comments the description "comment" to the photo "photoId". Returns the identifier of the comment.
// If an error occurs, returns 0.
func (db *appdbimpl) CommentPhoto(userId uint64, photoId uint64, comment string) (uint64, error) {
	// 1. Chequear si el usuario existe
	// 2. Chequear si la foto existe
	// 3. Si y solo si se cumplen las dos anteriores:
	// 3.1 Introducir una nueva relacion en Comments
	// 3.2 Incrementar el nComments en 1 de la foto

	// 1.
	if !db.UserExists(userId) {
		return 0, ErrUserSubjectNotExists
	}
	// 2.
	if !db.PhotoExists(photoId) {
		return 0, ErrPhotoNotExits
	}

	// 3.1
	res, err := db.c.Exec(`INSERT INTO Comments (user_id,photo_id,comment) VALUES (?, ?, ?)`,
		userId, photoId, comment)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	var idComment = uint64(lastInsertID)

	var nComments uint64
	if err := db.c.QueryRow("SELECT nComments FROM Photos where id = ?",
		photoId).Scan(&nComments); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
	}
	_, err = db.c.Exec(`UPDATE Photos SET nComments=? WHERE id=?`,
		nComments+1, photoId)
	if err != nil {
		return 0, err
	}
	return idComment, nil
}
