package erc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/x1rh/web3go/ethx/client"
)

type Adapter struct {
	*ERC20
	TokenAddress string
	ChainID      int64
	Client       *client.Client
}

func NewAdapter(ChainId int64, TokenAddress string, c *client.Client) (*Adapter, error) {
	erc20Obj, err := NewERC20(common.HexToAddress(TokenAddress), c.Client)
	if err != nil {
		return nil, errors.Wrap(err, "fail to initialize ERC20")
	}

	return &Adapter{
		ERC20:        erc20Obj,
		TokenAddress: TokenAddress,
		ChainID:      ChainId,
		Client:       c,
	}, nil
}

func MustNewAdapter(ChainId int64, tokenAddress string, c *client.Client) *Adapter {
	a, err := NewAdapter(ChainId, tokenAddress, c)
	if err != nil {
		panic(err)
	}
	return a
}
