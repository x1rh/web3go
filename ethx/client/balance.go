package client

import (
	"context"
	"github.com/x1rh/ethx/convertx"
	"math/big"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"
)

// Balance get current balance in wei
func (a *Adapter) Balance(ctx context.Context, walletAddress string) (*big.Int, error) {
	account := common.HexToAddress(walletAddress)
	balance, err := a.Client.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return balance, nil
}

// EtherBalance get current balance in eth
func (a *Adapter) EtherBalance(ctx context.Context, walletAddress string) (*decimal.Decimal, error) {
	balanceInWei, err := a.Balance(ctx, walletAddress)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return convertx.DivByDecimal(balanceInWei, 18)
}
