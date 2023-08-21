package horizon

import (
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/xdrbuild"
)

const (
	adminSeed = "SAMJKTZVW5UOHCDK5INYJNORF2HRKYI72M5XSZCBYAHQHR34FFR4Z6G4"
)

func formTransaction(op xdrbuild.Operation) (string, error) {
	passphrase := "TokenD Developer Network"
	builder := xdrbuild.NewBuilder(passphrase, 100)
	source, _ := keypair.Parse(adminSeed)
	transaction := builder.Transaction(source)
	transaction.Op(op)
	transaction.Sign(source.(*keypair.Full))
	return transaction.Marshal()
}
