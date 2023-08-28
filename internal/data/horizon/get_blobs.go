package horizon

import (
	"blob-service/internal/data"
	"encoding/json"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/go/keypair"
	"net/http"
	"strconv"
)

func GetBlobs(pageParams pgdb.OffsetPageParams, adminSeed string, getBlobsURL string) ([]data.BlobEntity, error) {
	source, _ := keypair.Parse(adminSeed)
	pageNumber := strconv.FormatUint(pageParams.PageNumber, 10)
	pageLimit := strconv.FormatUint(pageParams.Limit, 10)
	url := getBlobsURL + "?filter[owner]=" + source.Address() + "&page[number]=" + pageNumber + "&page[limit]=" + pageLimit + "&page[order]=" + pageParams.Order
	response, err := getAPIBlobs(url)
	if err != nil {
		return nil, err
	}
	return apiDataListToEntities(*response), nil
}

func getAPIBlobs(url string) (*apiDataList, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	dataList := new(apiDataList)
	err = json.NewDecoder(response.Body).Decode(dataList)
	if err != nil {
		return nil, err
	}
	return dataList, nil
}
