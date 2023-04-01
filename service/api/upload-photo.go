package api

import (
	"WASA/service/api/reqcontext"
	"WASA/service/database"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var url string
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var id uint64
	err = json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idPhoto, err := rt.db.UploadPhoto(url, id)
	if errors.Is(err, database.UserSubjectNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).WithField("user: ", id).Error("can't upload the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(idPhoto)

}
