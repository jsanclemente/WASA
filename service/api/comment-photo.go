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

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userId uint64
	err = json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment string
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nComments, err := rt.db.CommentPhoto(userId, photoId, comment)
	// User doesn't exists
	if errors.Is(err, database.UserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// Photo doesn't exists
	if errors.Is(err, database.ErrPhotoNotExits) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("user: ", userId).Error("can't comment the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nComments)

}
