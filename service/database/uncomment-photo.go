package database

import "database/sql"

// The user identified by "userId" uncomments the description "comment" to the photo "photoId". Returns the number of comments after the operation.
// If an error occurs, returns 0.
func (db *appdbimpl) UncommentPhoto(photoId uint64, commentId uint64) (uint64, error) {
	// 1. Chequear si la foto existe
	// 2. Chequear si el comentario existe
	// 3. Extraer el userId de la relación "Comments"
	// 4. Si y solo si se cumplen las condiciones 1 y 2:
	// 4.1 Eliminar la relacion en Comments
	// 4.2 Decrementear el nComments en 1 en la tabla Photos

	// 1.
	if !db.PhotoExists(photoId) {
		return 0, ErrPhotoNotExits
	}
	// 2.
	var idComment uint64
	var idUser uint64
	if err := db.c.QueryRow("SELECT user_id, comment_id FROM Comments WHERE comment_id = ?",
		commentId).Scan(&idUser, &idComment); err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrCommentNotExists
		}
	}

	// 4.1 User, photo and comment exists. Delete the comment
	_, err := db.c.Exec(`DELETE FROM Comments WHERE user_id=? AND comment_id=? AND photo_id=?`,
		idUser, commentId, photoId)
	if err != nil {
		return 0, err
	}

	// 4.2 Decrement the number of comments
	var nComments uint64
	if err := db.c.QueryRow("SELECT nComments FROM Photos where id = ?",
		photoId).Scan(&nComments); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
	}
	_, err = db.c.Exec(`UPDATE Photos SET nComments=? WHERE id=?`,
		nComments-1, photoId)
	if err != nil {
		return 0, err
	}

	return nComments - 1, nil
}
