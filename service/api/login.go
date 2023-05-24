package api

import (
	"WASA/service/api/reqcontext"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type LoginRequest struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserId uint64 `json:"userId"`
}

func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Hacer el login
	id, err := rt.db.Login(loginReq.Username)
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't create the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generar el token JWT con el userId

	token, err := generateToken(strconv.FormatUint(id, 10))
	if err != nil {
		// Manejo del error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Crear la respuesta con el token y el userId
	response := LoginResponse{
		Token:  token,
		UserId: id,
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

// Función para generar un token JWT firmado con el userId
func generateToken(userId string) (string, error) {
	// Crear un nuevo token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Configurar las reclamaciones (claims) del token
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Configurar la fecha de expiración del token

	// Firmar el token con una clave secreta
	secretKey := "9K7Ufvg$YmqP4e^u"

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
