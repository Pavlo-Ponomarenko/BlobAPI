package handlers

import (
	"blob-service/internal/service/requests"
	res "blob-service/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func CreateNewBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	q := BlobsQ(r)
	blob, err := q.SaveBlob(&request.Data)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	ape.Render(w, res.BlobResponse{Data: *blob})
}
