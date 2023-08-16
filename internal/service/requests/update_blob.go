package requests

import (
	res "blob-service/resources"
	"encoding/json"
	"net/http"
)

type UpdateBlobRequest struct {
	Id        string
	BlobModel UpdateBlobModel
}

type UpdateBlobModel struct {
	Data res.Blob `json:"data"`
}

func NewUpdateBlobRequest(r *http.Request) (*UpdateBlobRequest, error) {
	id, err := retrieveId(r)
	blobModel := new(UpdateBlobModel)
	err = json.NewDecoder(r.Body).Decode(blobModel)
	if err != nil {
		return nil, err
	}
	request := new(UpdateBlobRequest)
	request.Id = *id
	request.BlobModel = *blobModel
	return request, nil
}
