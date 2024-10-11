package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)


func GetTransactOpts(pk string, chainId *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get private key")
	}
	signTx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get transaction")
	}
	return signTx, nil
}


// BiggerGasPrice trick
func BiggerGasPrice(gasPrice *big.Int) *big.Int {
	gasPriceD := decimal.NewFromBigInt(gasPrice, 0)
	onePointTwo := decimal.NewFromFloat(1.2) 
	return gasPriceD.Mul(onePointTwo).BigInt()
}