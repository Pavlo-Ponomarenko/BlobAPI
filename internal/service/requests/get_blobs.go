package requests

import "net/http"

type GetBlobsRequest struct {
	Params map[string]string
}

func NewGetBlobsRequest(r *http.Request) *GetBlobsRequest {
	queryParams := r.URL.Query()
	params := make(map[string]string)
	pageLimit := queryParams.Get("page[limit]")
	if pageLimit != "" {
		params["limit"] = pageLimit
	}
	pageNumber := queryParams.Get("page[number]")
	if pageNumber != "" {
		params["number"] = pageNumber
	}
	pageOrder := queryParams.Get("page[order]")
	if pageOrder != "" {
		params["order"] = pageOrder
	}
	request := new(GetBlobsRequest)
	request.Params = params
	return request
}
