package erc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func (a *Adapter) Approve(
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

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(a.ChainID))
	if err != nil {
		return nil, err
	}
	txOpts.GasLimit = gasLimit
	txOpts.GasPrice = gasPrice

	return a.ERC20.Approve(txOpts, common.HexToAddress(spender), amount)
}
