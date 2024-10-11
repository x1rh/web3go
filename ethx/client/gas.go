package client

import (
	"context"
	"math/big"

	"github.com/x1rh/web3go/ethx/wallet"

	"github.com/ethereum/go-ethereum"
	"github.com/pkg/errors"
)

// EstimateGas
// get gas limit with call data
// fromAddress must not be empty
// toAddr could be empty
// ethValue could be nil or zero, unit is wei
func (c *Client) EstimateGas(ctx context.Context, fromAddr, toAddr any, data []byte, ethValue *big.Int) (uint64, error) {
	_from, err := wallet.MustAddress(fromAddr)
	if err != nil {
		return 0, errors.Wrap(err, "invalid from address")
	}

	_to, err := wallet.MustAddress(toAddr)
	if err != nil {
		return 0, errors.Wrap(err, "invalid to address")
	}

	return c.Client.EstimateGas(ctx, ethereum.CallMsg{
		From:     *_from,
		To:       _to,
		Gas:      0,
		GasPrice: big.NewInt(0),
		Value:    ethValue,
		Data:     data,
	})
}

// GasPrice get current gas price, unit is wei
func (c *Client) GasPrice(ctx context.Context) (*big.Int, error) {
	gasPrice, err := c.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	} else {
		return gasPrice, nil
	}
}

func (c *Client) EIP1559() {

}
