package requests

import (
	res "blob-service/resources"
	"encoding/json"
	"net/http"
)

type CreateBlobRequest struct {
	Data res.Blob `json:"data"`
}

func NewCreateBlobRequest(r *http.Request) (*CreateBlobRequest, error) {
	blob := new(CreateBlobRequest)
	if err := json.NewDecoder(r.Body).Decode(blob); err != nil {
		return nil, err
	}
	return blob, nil
}
