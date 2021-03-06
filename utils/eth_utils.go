package utils

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"

	ethComm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type EthSigner struct {
	PrivateKey *ecdsa.PrivateKey
	Address    ethComm.Address
}

func NewEthSigner(privateKey string) (*EthSigner, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("ApproveERC20, cannot decode private key")
	}
	address := crypto.PubkeyToAddress(priKey.PublicKey)
	return &EthSigner{
		PrivateKey: priKey,
		Address:    address,
	}, nil
}

func DeserializeTx(rawTx string) (*types.Transaction, error) {
	txData := ethComm.FromHex(rawTx)
	tx := &types.Transaction{}
	err := rlp.DecodeBytes(txData, tx)
	if err != nil {
		return nil, fmt.Errorf("deserializeTx: err: %s", err)
	}
	return tx, nil
}

func SerializeTx(tx *types.Transaction) (string, error) {
	bf := new(bytes.Buffer)
	err := rlp.Encode(bf, tx)
	if err != nil {
		return "", fmt.Errorf("signTx: encode signed tx err: %s", err)
	}
	signedRawTx := hexutil.Encode(bf.Bytes())
	return signedRawTx, nil
}
