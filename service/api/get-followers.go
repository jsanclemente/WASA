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

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var users []database.User
	users, err = rt.db.GetFollowers(userId)
	if errors.Is(err, database.ErrUserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("The user that starts the action does not exist"))
		if err != nil {
			return
		}
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("userId", userId).Error("can't get the followers")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}
