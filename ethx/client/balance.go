package client

import (
	"context"
	"math/big"

	"github.com/x1rh/web3go/ethx/convertx"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"
)

// Balance get current balance in wei
func (c *Client) Balance(ctx context.Context, walletAddress string) (*big.Int, error) {
	account := common.HexToAddress(walletAddress)
	balance, err := c.Client.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return balance, nil
}

// EtherBalance get current balance in eth
func (c *Client) EtherBalance(ctx context.Context, walletAddress string) (*decimal.Decimal, error) {
	balanceInWei, err := c.Balance(ctx, walletAddress)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return convertx.DivByDecimal(balanceInWei, 18)
}
