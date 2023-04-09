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

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The user id in the path is a 64-bit unsigned integer. Let's parse it.
	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// The followed id in the path is a 64-bit unsigned integer. Let's parse it.
	followedId, err := strconv.ParseUint(ps.ByName("followedId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nfollowers, err := rt.db.UnfollowUser(userId, followedId)
	if errors.Is(err, database.UserSubjectNotExists) {
		// The user (indicated by `id`) does not exist, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This user don't exists"))
		return
	}
	if errors.Is(err, database.UserPredicateNotExists) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The user you are trying to unfollow does not exist "))
		return
	}
	if errors.Is(err, database.ErrNotFollowing) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You can't unfollow a user that you are not following"))
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("userId", userId).Error("can't unfollow the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nfollowers)
}
