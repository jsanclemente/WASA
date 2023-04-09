package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SetNameRequest struct {
	Username string `json:"username"`
}

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var request SetNameRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oldUsername, err := rt.db.SetMyUserName(userId, request.Username)
	if errors.Is(err, database.UserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The user that starts the action does not exist"))
		return
	}
	if errors.Is(err, database.ErrUsernameAlreadyRegistered) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("This username is already used by another user"))
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("no se ha podido cambiar username: ", request.Username).Error("can't unban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(oldUsername)
}
