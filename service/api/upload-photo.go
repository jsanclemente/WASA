package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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

		err = r.ParseMultipartForm(32 << 20) // 32 MB
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		idString := r.FormValue("userId")

		expectedUserId := idString
		// Comparar el userId del token con el valor esperado
		if userId != expectedUserId {
			// El userId en el token no coincide con el valor esperado
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// El token es válido y el userId es correcto

		id, _ := strconv.ParseUint(idString, 10, 64)
		image, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer image.Close()

		// Leer el archivo en memoria
		imageBytes, err := io.ReadAll(image)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idPhoto, err := rt.db.UploadPhoto(imageBytes, id)
		if errors.Is(err, database.ErrUserSubjectNotExists) {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("The user that uploads the photo does not exist"))
			if err != nil {
				return
			}
			return
		}
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).WithField("user: ", userId).Error("can't upload the photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(idPhoto)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
