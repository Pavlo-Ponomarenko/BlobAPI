package handlers

import (
	"blob-service/internal/data"
	"blob-service/internal/service/requests"
	res "blob-service/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func UpdateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewUpdateBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	q := BlobsQ(r)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	entity, err := q.UpdateBlob(request.Id, data.BlobToEntity(&request.BlobModel.Data))
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	blob := data.EntityToBlob(entity)
	ape.Render(w, res.BlobResponse{Data: *blob})
}
