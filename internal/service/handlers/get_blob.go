package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	res "blob-service/resources"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

func GetBlobById(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobByIdRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q, err := data.CreateNewBlobsQ()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer q.Close()
	model, err := q.GetBlobById(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, res.BlobModelResponse{Data: *model})
}
