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

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nlikes, err := rt.db.LikePhoto(userId, photoId)
	if errors.Is(err, database.UserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The user that starts the action does not exist"))
		return
	}
	if errors.Is(err, database.ErrPhotoNotExits) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("You can't like a photo that does not exist"))
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("photoId", photoId).Error("can't like the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nlikes)
}
