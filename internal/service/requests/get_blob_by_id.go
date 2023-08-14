package requests

import "net/http"

type GetBlobByIdRequest struct {
	ID string
}

func NewGetBlobByIdRequest(r *http.Request) (*GetBlobByIdRequest, error) {
	id, err := retrieveId(r)
	if err != nil {
		return nil, err
	}
	request := new(GetBlobByIdRequest)
	request.ID = *id
	return request, nil
}
