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

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Obtener el userId del token
		userId, ok := claims["userId"].(string)
		if !ok {
			// El userId no es una cadena válida en el token
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// The user id in the path is a 64-bit unsigned integer. Let's parse it.
		expectedUserId := ps.ByName("userId")
		// Comparar el userId del token con el valor esperado
		if userId != expectedUserId {
			// El userId en el token no coincide con el valor esperado
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// El token es válido y el userId es correcto
		photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
		if err != nil {
			// The value was not uint64, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			// The value was not uint64, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		nlikes, err := rt.db.UnlikePhoto(id, photoId)
		if errors.Is(err, database.ErrUserSubjectNotExists) {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("The user that unlikes the photo does not exist"))
			if err != nil {
				return
			}
			return
		}
		if errors.Is(err, database.ErrNotHisLike) {
			w.WriteHeader(http.StatusForbidden)
			_, err := w.Write([]byte("You can't remove a like that is not yours"))
			if err != nil {
				return
			}
			return
		}
		if errors.Is(err, database.ErrPhotoNotExits) {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("You can remove the like of a photo that does not exist"))
			if err != nil {
				return
			}
			return
		} else if err != nil {
			ctx.Logger.WithError(err).WithField("photoId", photoId).Error("can't unlike the photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(nlikes)
	}
}
