package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := r.URL.Query().Get("username")
	query := r.URL.Query().Get("query")
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
	users, err := rt.db.SearchUser(username, query)
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

}
