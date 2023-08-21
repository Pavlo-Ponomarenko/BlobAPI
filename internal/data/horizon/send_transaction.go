package horizon

import (
	"blob-service/internal/data"
	"bytes"
	"encoding/json"
	"errors"
	"gitlab.com/tokend/go/xdr"
	"io"
	"net/http"
)

const (
	apiUrl = "http://localhost:8000/_/api/"
)

type createBlobResult struct {
	Data map[string]interface{} `json:"data"`
}

func sendTransaction(endpoint string, jsonData []byte) (*data.BlobEntity, error) {
	httpRequest, err := http.NewRequest("POST", apiUrl+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return getBlobFromResult(body)
}

func getBlobFromResult(result []byte) (*data.BlobEntity, error) {
	var transactionResult createBlobResult
	err := json.Unmarshal(result, &transactionResult)
	if err != nil {
		return nil, err
	}
	dict := transactionResult.Data["attributes"].(map[string]interface{})
	encodedStatus := dict["result_xdr"]
	var responseStatus xdr.TransactionResult
	err = xdr.SafeUnmarshalBase64(encodedStatus.(string), &responseStatus)
	if responseStatus.Result.Code != xdr.TransactionResultCodeTxSuccess {
		return nil, errors.New("")
	}
	encodedMeta := dict["result_meta_xdr"]
	var resultMeta xdr.TransactionMeta
	err = xdr.SafeUnmarshalBase64(encodedMeta.(string), &resultMeta)
	if err != nil {
		return nil, err
	}
	decodedBlob := (*resultMeta.Operations)[0].Changes[0].Created.Data.Data.Value
	createdBlob := new(data.BlobEntity)
	err = json.Unmarshal([]byte(decodedBlob), createdBlob)
	if err != nil {
		return nil, err
	}
	return createdBlob, nil
}
