package requests

import (
	"blob-service/internal/service/handlers"
	res "blob-service/resources"
	"encoding/json"
	//	"fmt"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func retrieveId(r *http.Request) (*string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("")
	}
	return &id, nil
}

func GetBlobById(w http.ResponseWriter, r *http.Request) {
	id, err := retrieveId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	blob, err := handlers.GetBlobById(*id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, res.BlobModelResponse{Data: *blob})
}

func CreateNewBlob(w http.ResponseWriter, r *http.Request) {
	newBlob := new(res.Blob)
	err := json.NewDecoder(r.Body).Decode(newBlob)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handlers.SaveBlob(newBlob)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	id, err := retrieveId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handlers.DeleteBlob(*id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetPageOfBlobs(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	params := make(map[string]string)
	pageLimit := queryParams.Get("page[limit]")
	if pageLimit != "" {
		params["limit"] = pageLimit
	}
	pageNumber := queryParams.Get("page[number]")
	if pageNumber != "" {
		params["number"] = pageNumber
	}
	pageOrder := queryParams.Get("page[order]")
	if pageOrder != "" {
		params["order"] = pageOrder
	}
	response, err := handlers.GetBlobs(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, response)
}

func UpdateBlob(w http.ResponseWriter, r *http.Request) {
	id, err := retrieveId(r)
	newBlob := new(res.Blob)
	err = json.NewDecoder(r.Body).Decode(newBlob)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !handlers.IdIsPresent(*id) {
		w.WriteHeader(http.StatusNotFound)
	}
	if err := handlers.UpdateBlob(*id, newBlob); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
