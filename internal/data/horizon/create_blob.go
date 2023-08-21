package horizon

import (
	"blob-service/internal/data"
	"encoding/json"
	"gitlab.com/tokend/go/xdrbuild"
)

type createBlobRequest struct {
	Tx            string `json:"tx"`
	WaitForIngest bool   `json:"wait_for_ingest"`
	WaitForResult bool   `json:"wait_for_result"`
}

type blobToJSON struct {
	data.BlobEntity
}

func (b blobToJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BlobEntity)
}

func CreateBlob(entity *data.BlobEntity) (*data.BlobEntity, error) {
	createDataOp := xdrbuild.CreateData{Type: 1, Value: blobToJSON{*entity}}
	transaction, err := formTransaction(createDataOp)
	request := createBlobRequest{
		Tx:            transaction,
		WaitForIngest: false,
		WaitForResult: true,
	}
	jsonRequest, _ := json.Marshal(request)
	result, err := sendTransaction("v3/transactions", jsonRequest)
	if err != nil {
		return nil, err
	}
	return result, nil
}
