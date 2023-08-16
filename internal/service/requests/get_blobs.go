package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"strconv"
)

type GetBlobsRequest struct {
	Params pgdb.OffsetPageParams
}

func NewGetBlobsRequest(r *http.Request) (*GetBlobsRequest, error) {
	queryParams := r.URL.Query()
	params := pgdb.OffsetPageParams{}
	pageLimit := queryParams.Get("page[limit]")
	if pageLimit != "" {
		value, err := strconv.ParseUint(pageLimit, 10, 64)
		if err != nil {
			return nil, err
		}
		params.Limit = value
	}
	pageNumber := queryParams.Get("page[number]")
	if pageNumber != "" {
		value, err := strconv.ParseUint(pageNumber, 10, 64)
		if err != nil {
			return nil, err
		}
		params.PageNumber = value
	}
	pageOrder := queryParams.Get("page[order]")
	if pageOrder != "" {
		if pageOrder == "asc" || pageOrder == "desc" {
			params.Order = pageOrder
		} else {
			return nil, errors.New("")
		}
	}
	request := new(GetBlobsRequest)
	request.Params = params
	return request, nil
}
