package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	res "blob-service/resources"
	"gitlab.com/distributed_lab/ape"
	"net/http"
)

func GetPageOfBlobs(w http.ResponseWriter, r *http.Request) {
	request := requests.NewGetBlobsRequest(r)
	q, err := data.CreateNewBlobsQ()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer q.Close()
	response, err := q.GetBlobs(request.Params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, res.BlobModelListResponse{Data: response})
}
