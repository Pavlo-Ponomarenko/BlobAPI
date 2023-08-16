package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	res "blob-service/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetBlobById(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	q := BlobsQ(r)
	blobEntity, err := q.GetBlobById(request.ID)
	if err != nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	blob := data.EntityToBlob(blobEntity)
	ape.Render(w, res.BlobResponse{Data: *blob})
}
