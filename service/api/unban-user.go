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

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The user id in the path is a 64-bit unsigned integer. Let's parse it.
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
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

	nbanned, err := rt.db.UnbanUser(userId, bannedId)
	if errors.Is(err, database.ErrUserSubjectNotExists) {
		// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("The user that starts the action does not exist"))
		if err != nil {
			return
		}
		return
	}
	if errors.Is(err, database.ErrNotBanned) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("You can't unban a user that is not banned"))
		if err != nil {
			return
		}
		return
	}
	if errors.Is(err, database.ErrUserPredicateNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("The user that do you are trying to unban does not exist"))
		if err != nil {
			return
		}
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("bannedId", bannedId).Error("can't unban the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nbanned)
}
