package handlers

import (
	"blob-service/internal/service/requests"
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
	if !q.IdIsPresent(request.Id) {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if err := q.UpdateBlob(request.Id, &request.BlobModel.Data); err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusOK)
}
