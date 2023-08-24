package horizon

import (
	"blob-service/internal/data"
	"gitlab.com/tokend/go/xdrbuild"
	"strconv"
)

func UpdateBlob(id string, entity *data.BlobEntity) (*data.BlobEntity, error) {
	dataId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	updateDataOp := xdrbuild.UpdateData{ID: dataId, Value: blobToJSON{*entity}}
	transaction, _ := formTransaction(updateDataOp)
	jsonRequest := formJsonRequest(transaction)
	result, err := sendTransaction("v3/transactions", jsonRequest)
	if err != nil {
		return nil, err
	}
	return getBlobFromResult(result, updateBlobTransaction)
}
