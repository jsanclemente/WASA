package database

import (
	"database/sql"
	"fmt"
)

// The user identified by "userId" comments the description "comment" to the photo "photoId". Returns the number of comments after the operation.
// If an error occurs, returns 0.
func (db *appdbimpl) CommentPhoto(userId uint64, photoId uint64, comment string) (uint64, error) {
	// 1. Chequear si el usuario existe
	// 2. Chequear si la foto existe
	// 3. Si y solo si se cumplen las dos anteriores:
	// 3.1 Introducir una nueva relacion en Comments
	// 3.2 Incrementar el nComments en 1 de la foto

	// 1.
	if !db.UserExists(userId) {
		return 0, UserSubjectNotExists
	}
	// 2.
	if !db.PhotoExists(photoId) {
		return 0, ErrPhotoNotExits
	}

	// 3.1
	_, err := db.c.Exec(`INSERT INTO Comments (user_id,photo_id,comment) VALUES (?, ?, ?)`,
		userId, photoId, comment)
	if err != nil {
		fmt.Print("Error en el Insert")
		return 0, err
	}

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
		fmt.Print("Error en el Update")
		return 0, err
	}
	return nComments + 1, nil
}
