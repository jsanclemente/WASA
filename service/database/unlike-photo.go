package database

import "database/sql"

// The user "userId" unlikes the photo identified by "photoId". Returns the number of likes after the operation.
// If an error occurs, returns 0.
func (db *appdbimpl) UnlikePhoto(userId uint64, photoId uint64) (uint64, error) {

	// 1. Comprobar si existe el userId
	// 2. Comprobar si existe el photoId
	// 3. Eliminar la relacion en la tabla "Likes"
	// 4. Decrementar el numero de likes

	// 1. Check if the user exists
	if !db.UserExists(userId) {
		return 0, UserSubjectNotExists
	}
	// 2. Check if the photo exists
	var nLikes uint64
	if err := db.c.QueryRow("SELECT nLikes FROM Photos where id = ?",
		photoId).Scan(&nLikes); err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrPhotoNotExits
		}
	}
	// 3. Both userId and photoId, exists, delete the row
	_, err := db.c.Exec(`DELETE FROM Likes WHERE user_id=? AND photo_id=?`,
		userId, photoId)
	if err != nil {
		return 0, err
	}
	// 4. Decrement the value of nLikes of that photo
	_, err = db.c.Exec(`UPDATE Photos SET nLikes=? WHERE id=?`,
		nLikes-1, photoId)
	if err != nil {
		return 0, err
	}
	return nLikes - 1, nil
}
