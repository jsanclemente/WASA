package database

// userName posts the image "url". Returns the id of the photo posted. If something is wrong return the value 0.
func (db *appdbimpl) UploadPhoto(image []byte, id uint64) (uint64, error) {

	// 1. Insertar la foto en la tabla "Photos". Inicialmente 0 likes y 0 comments
	// 2. Insertar en la tabla Posts que el usuario id postea la foto idPhoto
	// 3. Incrementar el numero de posts de User en User + 1

	//Check if the user exists
	if !db.UserExists(id) {
		return 0, ErrUserSubjectNotExists
	}

	var idPhoto uint64
	// 1.
	res, err := db.c.Exec(`INSERT INTO Photos (id, nLikes, nComments, imageData) VALUES (NULL, ?, ?, ?)`,
		0, 0, image)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	idPhoto = uint64(lastInsertID)

	res, err = db.c.Exec(`INSERT INTO Posts (user_id,photo_id,date) VALUES (?, ?, datetime('now','localtime'))`,
		id, idPhoto)
	if err != nil {
		return 0, err
	}

	// Increment number of posts
	var nPosts uint64
	if err := db.c.QueryRow("SELECT nPosts FROM Users where id = ?", id).
		Scan(&nPosts); err != nil {
		return 0, err
	}
	nPosts = nPosts + 1
	_, err = db.c.Exec(`UPDATE Users SET nposts=? WHERE id=?`, nPosts, id)
	if err != nil {
		return 0, err
	}

	return idPhoto, nil
}
