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

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

		id, err := strconv.ParseUint(userId, 10, 64)
		if err != nil {
			// The value was not uint64, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// The followed id in the path is a 64-bit unsigned integer. Let's parse it.
		bannedId, err := strconv.ParseUint(ps.ByName("bannedId"), 10, 64)
		if err != nil {
			// The value was not uint64, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		nbanned, err := rt.db.BanUser(id, bannedId)
		if errors.Is(err, database.ErrUserSubjectNotExists) {
			// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("The user that do the ban does not exist"))
			if err != nil {
				return
			}
			return
		}
		if errors.Is(err, database.ErrAlreadyBanned) {
			w.WriteHeader(http.StatusForbidden)
			_, err := w.Write([]byte("The user that you are trying to ban is already bananed"))
			if err != nil {
				return
			}
			return
		}
		if errors.Is(err, database.ErrUserPredicateNotExists) {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("The user that you are trying to ban does not exist"))
			if err != nil {
				return
			}
			return
		}
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).WithField("bannedId", bannedId).Error("can't ban the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(nbanned)
	} else {
		// El token no es válido
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
