package horizon

import (
	"blob-service/internal/data"
	"gitlab.com/tokend/go/xdrbuild"
)

func CreateBlob(entity *data.BlobEntity) (*data.BlobEntity, error) {
	createDataOp := xdrbuild.CreateData{Type: 1, Value: blobToJSON{*entity}}
	transaction, err := formTransaction(createDataOp)
	jsonRequest := formJsonRequest(transaction)
	result, err := sendTransaction("v3/transactions", jsonRequest)
	if err != nil {
		return nil, err
	}
	return getBlobFromResult(result, createBlobTransaction)
}
