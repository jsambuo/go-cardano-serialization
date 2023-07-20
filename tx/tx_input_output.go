package tx

import (
	"encoding/hex"

	"github.com/fxamacker/cbor/v2"
	"github.com/jsambuo/go-cardano-serialization/address"
)

type TxInput struct {
	cbor.Marshaler

	TxHash []byte
	Index  uint16
	Amount uint
}

type Amount struct {
	Unit     string
	Quantity string
}
type Amounts struct {
	Amount []Amount
}
type TxInputWithNFTS struct {
	cbor.Marshaler

	TxHash []byte
	Index  uint16
	Amount Amounts
}

func NewTxInputWithNFTS(txHash string, txIx uint16, amount Amounts) *TxInputWithNFTS {
	hash, _ := hex.DecodeString(txHash)

	return &TxInputWithNFTS{
		TxHash: hash,
		Index:  txIx,
		Amount: amount,
	}
}

// NewTxInput creates and returns a *TxInput from Transaction Hash(Hex Encoded), Transaction Index and Amount.
func NewTxInput(txHash string, txIx uint16, amount uint) *TxInput {
	hash, _ := hex.DecodeString(txHash)

	return &TxInput{
		TxHash: hash,
		Index:  txIx,
		Amount: amount,
	}
}

func (txI *TxInput) MarshalCBOR() ([]byte, error) {
	type arrayInput struct {
		_      struct{} `cbor:",toarray"`
		TxHash []byte
		Index  uint16
	}
	input := arrayInput{
		TxHash: txI.TxHash,
		Index:  txI.Index,
	}
	return cbor.Marshal(input)
}

type TxOutput struct {
	_       struct{} `cbor:",toarray"`
	Address address.Address
	Amount  uint
}

func NewTxOutput(addr address.Address, amount uint) *TxOutput {
	return &TxOutput{
		Address: addr,
		Amount:  amount,
	}
}
