package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	"net/http"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q, err := data.CreateNewBlobsQ()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer q.Close()
	err = q.DeleteBlob(request.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
