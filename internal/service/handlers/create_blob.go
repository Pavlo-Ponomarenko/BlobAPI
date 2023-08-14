package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	"net/http"
)

func CreateNewBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q, err := data.CreateNewBlobsQ()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer q.Close()
	err = q.SaveBlob(&request.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
