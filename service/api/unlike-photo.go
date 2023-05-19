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

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	nlikes, err := rt.db.UnlikePhoto(userId, photoId)
	if errors.Is(err, database.ErrUserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("The user that unlikes the photo does not exist"))
		if err != nil {
			return
		}
		return
	}
	if errors.Is(err, database.ErrNotHisLike) {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("You can't remove a like that is not yours"))
		if err != nil {
			return
		}
		return
	}
	if errors.Is(err, database.ErrPhotoNotExits) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("You can remove the like of a photo that does not exist"))
		if err != nil {
			return
		}
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("photoId", photoId).Error("can't unlike the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nlikes)
}
