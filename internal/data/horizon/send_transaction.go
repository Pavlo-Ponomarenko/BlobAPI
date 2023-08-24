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
	sendTransactionURL = "http://localhost:8000/_/api/"
)

type transactionPostRequest struct {
	Tx            string `json:"tx"`
	WaitForIngest bool   `json:"wait_for_ingest"`
	WaitForResult bool   `json:"wait_for_result"`
}

type blobTransactionResult struct {
	Data map[string]interface{} `json:"data"`
}

const (
	createBlobTransaction = 0
	updateBlobTransaction = 1
)

type transactionType int

func sendTransaction(endpoint string, jsonData []byte) ([]byte, error) {
	httpRequest, err := http.NewRequest("POST", sendTransactionURL+endpoint, bytes.NewBuffer(jsonData))
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
	return body, nil
}

func getBlobFromResult(result []byte, resultType transactionType) (*data.BlobEntity, error) {
	dict := getAttributes(result)
	responseInfo := getResultXDR(dict)
	if responseInfo.Result.Code != xdr.TransactionResultCodeTxSuccess {
		return nil, errors.New("")
	}
	resultMeta := getResultMetaXDR(dict)
	var decodedBlob xdr.Longstring
	switch resultType {
	case createBlobTransaction:
		decodedBlob = (*resultMeta.Operations)[0].Changes[0].Created.Data.Data.Value
	case updateBlobTransaction:
		decodedBlob = (*resultMeta.Operations)[0].Changes[0].Updated.Data.Data.Value
	}
	createdBlob := new(data.BlobEntity)
	err := json.Unmarshal([]byte(decodedBlob), createdBlob)
	if err != nil {
		return nil, err
	}
	return createdBlob, nil
}

func getAttributes(result []byte) map[string]interface{} {
	var transactionResult blobTransactionResult
	_ = json.Unmarshal(result, &transactionResult)
	return transactionResult.Data["attributes"].(map[string]interface{})
}

func getResultXDR(attributes map[string]interface{}) xdr.TransactionResult {
	encodedStatus := attributes["result_xdr"]
	var responseStatus xdr.TransactionResult
	_ = xdr.SafeUnmarshalBase64(encodedStatus.(string), &responseStatus)
	return responseStatus
}

func getResultMetaXDR(attributes map[string]interface{}) xdr.TransactionMeta {
	encodedMeta := attributes["result_meta_xdr"]
	var resultMeta xdr.TransactionMeta
	_ = xdr.SafeUnmarshalBase64(encodedMeta.(string), &resultMeta)
	return resultMeta
}
