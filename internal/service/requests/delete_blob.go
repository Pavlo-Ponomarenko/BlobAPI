package requests

import "net/http"

type DeleteBlobRequest struct {
	Id string
}

func NewDeleteBlobRequest(r *http.Request) (*DeleteBlobRequest, error) {
	id, err := retrieveId(r)
	if err != nil {
		return nil, err
	}
	request := new(DeleteBlobRequest)
	request.Id = *id
	return request, nil
}
