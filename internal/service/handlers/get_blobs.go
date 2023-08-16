package handlers

import (
	"blob-service/internal/service/requests"
	res "blob-service/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetPageOfBlobs(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobsRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
	}
	q := BlobsQ(r)
	response, err := q.GetBlobs(request.Params)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	ape.Render(w, res.BlobListResponse{Data: response})
}
