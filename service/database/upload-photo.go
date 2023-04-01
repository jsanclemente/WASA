package database

import (
	"fmt"
	"time"
)

// userName posts the image "url". Returns the id of the photo posted. If something is wrong return the value 0.
func (db *appdbimpl) UploadPhoto(url string, id uint64) (uint64, error) {

	// 1. Insertar la foto en la tabla "Photos". Inicialmente 0 likes y 0 comments
	// 2. Insertar en la tabla Posts que el usuario id postea la foto idPhoto
	// 3. Incrementar el numero de posts de User en User + 1

	//Ceck if the user exists
	if !db.UserExists(id) {
		return 0, UserSubjectNotExists
	}

	var idPhoto uint64
	// 1.
	res, err := db.c.Exec(`INSERT INTO Photos (id, nLikes, nComments, url) VALUES (NULL, ?, ?, ?)`,
		0, 0, url)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	idPhoto = uint64(lastInsertID)

	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	res, err = db.c.Exec(`INSERT INTO Posts (user_id,photo_id,date) VALUES (?, ?, ?)`,
		id, idPhoto, date)
	if err != nil {
		return 0, err
	}
	return idPhoto, nil
}
