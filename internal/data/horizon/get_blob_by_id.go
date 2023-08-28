package horizon

import (
	"blob-service/internal/data"
	"encoding/json"
	"errors"
	"net/http"
)

func GetBlobById(id string, getBlobsURL string) (*data.BlobEntity, error) {
	url := getBlobsURL + "/" + id
	coreData, err := getAPIBlobById(url)
	if err != nil {
		return nil, err
	}
	result := apiDataToEntity(*coreData)
	return result, nil
}

func getAPIBlobById(url string) (*apiData, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("blob is not found")
	}
	coreData := new(apiDataById)
	err = json.NewDecoder(response.Body).Decode(coreData)
	if err != nil {
		return nil, err
	}
	return &coreData.Data, nil
}
