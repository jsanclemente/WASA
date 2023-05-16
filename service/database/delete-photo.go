package database

// Deletes the photo idPhoto. Returns the id of the deleted photo
// If an error occurs, returns 0
func (db *appdbimpl) DeletePhoto(idUser uint64, idPhoto uint64) (uint64, error) {

	// 1. Eliminar la foto de la tabla Photos, por tanto se elimina de Posts,Likes,Comments tambien
	// 2. Hacer el update de la tabla User de nPosts

	if !db.UserExists(idUser) {
		return 0, ErrUserSubjectNotExists
	}

	if !db.PhotoExists(idPhoto) {
		return 0, ErrPhotoNotExits
	}

	if !db.OwnsPhoto(idUser, idPhoto) {
		return 0, ErrNotHisPhoto
	}

	// Esta sentencia borra en la tabla Posts, Likes, Comments gracias al ON DELETE CASCADE
	_, err := db.c.Exec(`DELETE FROM Photos WHERE id=?`, idPhoto)
	if err != nil {
		return 0, err
	}

	var nPosts uint64
	if err := db.c.QueryRow("SELECT nPosts FROM Users where id = ?", idUser).Scan(&nPosts); err != nil {
		return 0, err
	}
	nPosts = nPosts - 1
	_, err = db.c.Exec(`UPDATE Users SET nposts=? WHERE id=?`, nPosts, idUser)
	if err != nil {
		return 0, err
	}
	return idPhoto, nil
}
