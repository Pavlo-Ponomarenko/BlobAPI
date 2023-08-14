package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	"net/http"
)

func UpdateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateBlobRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q, err := data.CreateNewBlobsQ()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer q.Close()
	if !q.IdIsPresent(request.Id) {
		w.WriteHeader(http.StatusNotFound)
	}
	if err := q.UpdateBlob(request.Id, &request.Data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
