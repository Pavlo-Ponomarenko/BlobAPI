package requests

import (
	res "blob-service/resources"
	"encoding/json"
	"net/http"
)

type CreateBlobRequest struct {
	Data res.Blob
}

func NewCreateBlobRequest(r *http.Request) (*CreateBlobRequest, error) {
	blob := new(res.Blob)
	if err := json.NewDecoder(r.Body).Decode(blob); err != nil {
		return nil, err
	}
	request := new(CreateBlobRequest)
	request.Data = *blob
	return request, nil
}
