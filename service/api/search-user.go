package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

		// Get the id on the body
		idString := r.URL.Query().Get("id")
		expectedUserId := idString
		// Comparar el userId del token con el valor esperado
		if userId != expectedUserId {
			// El userId en el token no coincide con el valor esperado
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// El token es válido y el userId es correcto
		username := r.URL.Query().Get("username")
		query := r.URL.Query().Get("query")

		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			// Manejar el error de conversión
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// If username is empty, the request is wrong
		if username == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var users []database.User
		users, err = rt.db.SearchUser(username, query, id)
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).WithField("error searching", username).Error("can't search this user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(users)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
