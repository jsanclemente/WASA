package database

import (
	"fmt"
)

// The user "userId" likes the photo identified by "photoId". Returns the number of likes after the operation
func (db *appdbimpl) LikePhoto(userId uint64, photoId uint64) (uint64, error) {

	// 1. Chequear si el usuario existe
	// 2. Chequear si la foto existe
	// 3. Si y solo si se cumplen las dos anteriores:
	// Introducir una nueva relacion en Likes
	// Incrementar el nLikes en 1 de la foto

	if !db.UserExists(userId) {
		return 0, UserSubjectNotExists
	}
	// 2. User exists, now we have to check whether the photo does
	var nLikes uint64
	if err := db.c.QueryRow("SELECT nLikes FROM Photos where id = ?",
		photoId).Scan(&nLikes); err != nil {
		return 0, ErrPhotoNotExits
	}
	// 3. Both userId and photoId, exists
	// Check the "userId" doesn't like a photo twice
	if db.AlreadyLiked(userId, photoId) {
		return 0, ErrPhotoAlreadyLiked
	}

	_, err := db.c.Exec(`INSERT INTO Likes (user_id,photo_id) VALUES (?, ?)`,
		userId, photoId)
	if err != nil {
		fmt.Print("error en el insert en likes\n")
		return 0, err
	}
	// Update the value of nLikes of that photo
	_, err = db.c.Exec(`UPDATE Photos SET nLikes=? WHERE id=?`,
		nLikes+1, photoId)
	if err != nil {
		return 0, err
	}
	fmt.Print("Llego al update \n")
	return nLikes + 1, nil
}
