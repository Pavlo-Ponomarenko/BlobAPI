package horizon

import (
	"encoding/json"
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/xdrbuild"
)

var cashedCoreInfo coreInfoCash

func formTransaction(op xdrbuild.Operation, adminSeed string, coreInfoURL string) (string, error) {
	info, err := cashedCoreInfo.getInfo(coreInfoURL)
	if err != nil {
		return "", err
	}
	builder := xdrbuild.NewBuilder(info.Attributes.NetworkPassphrase, info.Attributes.TxExpirationPeriod)
	source, _ := keypair.Parse(adminSeed)
	transaction := builder.Transaction(source)
	transaction.Op(op)
	transaction.Sign(source.(*keypair.Full))
	return transaction.Marshal()
}

func formJsonRequest(transaction string) []byte {
	request := transactionPostRequest{
		Tx:            transaction,
		WaitForIngest: false,
		WaitForResult: true,
	}
	jsonRequest, _ := json.Marshal(request)
	return jsonRequest
}
