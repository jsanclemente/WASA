package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Verificar y validar el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar el método de firma y obtener la clave secreta
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Retorna la clave secreta
		return []byte("9K7Ufvg$YmqP4e^u"), nil
	})

	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Token is valid

	query := r.URL.Query()
	idString := query.Get("userId")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// The user id in the path is a 64-bit unsigned integer. Let's parse photoId.
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postId, err := rt.db.DeletePhoto(id, photoId)
	if errors.Is(err, database.ErrUserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("The user that deletes the photo does not exist"))
		if err != nil {
			return
		}
		return
	}
	if errors.Is(err, database.ErrNotHisPhoto) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("You can't delete a photo you have not posted"))
		if err != nil {
			return
		}
		return
	}
	if errors.Is(err, database.ErrPhotoNotExits) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("This photo does not exist"))
		if err != nil {
			return
		}
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("photoId", photoId).Error("can't delete the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(postId)
}
