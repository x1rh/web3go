package client

import (
	"context"
	"math/big"

	"github.com/x1rh/web3go/ethx/castx"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// Balance get current balance in wei
func (c *Client) Balance(ctx context.Context, walletAddress any) (*big.Int, error) {
	account, err := castx.ToAddress(walletAddress)
	if err != nil {
		return nil, err
	}
	balance, err := c.Client.BalanceAt(ctx, *account, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return balance, nil
}

// EtherBalance get current balance in eth
func (c *Client) EtherBalance(ctx context.Context, walletAddress any) (*decimal.Decimal, error) {
	balanceInWei, err := c.Balance(ctx, walletAddress)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return castx.DivByDecimal(balanceInWei, 18)
}
