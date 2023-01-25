package database

import "database/sql"

// The user identified by "userId" comments the description "comment" to the photo "photoId". Returns the number of comments after the operation.
// If an error occurs, returns 0.
func (db *appdbimpl) CommentPhoto(userId uint64, photoId uint64, comment string) (uint64, error){
	// 1. Chequear si el usuario existe
	// 2. Chequear si la foto existe
	// 3. Si y solo si se cumplen las dos anteriores:
		// Introducir una nueva relacion en Comments
		// Incrementar el nComments en 1 de la foto 

}