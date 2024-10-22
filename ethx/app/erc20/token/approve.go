package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (t *Token) Approve(
	pk string,
	spender string,
	amount *big.Int,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(t.ChainId)))
	if err != nil {
		return nil, err
	}
	txOpts.GasLimit = gasLimit
	txOpts.GasPrice = gasPrice

	return t.ERC20.Approve(txOpts, common.HexToAddress(spender), amount)
}
