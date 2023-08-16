package handlers

import (
	"blob-service/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	q := BlobsQ(r)
	err = q.DeleteBlob(request.Id)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
