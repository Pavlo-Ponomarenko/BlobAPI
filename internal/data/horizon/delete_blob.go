package horizon

import (
	"errors"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/go/xdrbuild"
	"strconv"
)

func DeleteBlob(id string) error {
	dataId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	removeDataOp := xdrbuild.RemoveData{ID: dataId}
	transaction, err := formTransaction(removeDataOp)
	jsonRequest := formJsonRequest(transaction)
	result, err := sendTransaction("v3/transactions", jsonRequest)
	if err != nil {
		return err
	}
	dict := getAttributes(result)
	responseInfo := getResultXDR(dict)
	if responseInfo.Result.Code != xdr.TransactionResultCodeTxSuccess {
		return errors.New("remove operation failed")
	}
	return nil
}
