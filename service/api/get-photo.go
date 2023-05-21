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

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		// The value was not uint64, reject the action indicating an error on the client side.
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	var photo database.Photo
	photo, err = rt.db.GetPhoto(photoId)
	if errors.Is(err, database.ErrUserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("This photo does not exist"))
		if err != nil {
			return
		}
		return
	}

	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("no se ha podido obtener la foto: ", photoId).Error("error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)

}