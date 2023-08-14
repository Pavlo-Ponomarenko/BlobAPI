package requests

import (
	res "blob-service/resources"
	"encoding/json"
	"net/http"
)

type UpdateBlobRequest struct {
	Id   string
	Data res.Blob
}

func NewUpdateBlobRequest(r *http.Request) (*UpdateBlobRequest, error) {
	id, err := retrieveId(r)
	blob := new(res.Blob)
	err = json.NewDecoder(r.Body).Decode(blob)
	if err != nil {
		return nil, err
	}
	request := new(UpdateBlobRequest)
	request.Id = *id
	request.Data = *blob
	return request, nil
}
