package token

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/x1rh/web3go/ethx/client"
)

type TokenInfo struct {
	ChainId  int
	Address  string
	Metadata *TokenMetadata
}

type TokenMetadata struct {
	Name     string
	Symbol   string
	Decimals uint64
}

type Token struct {
	*ERC20

	Client *client.Client

	ChainId int
	Address string

	// we will initialize Metadata after calling `TokenMetadata()` at first time
	// read the data from block chain and cached it
	Metadata *TokenMetadata
}

func New(chainId int, address string, c *client.Client) (*Token, error) {
	token, err := NewERC20(common.HexToAddress(address), c.Client)
	if err != nil {
		return nil, errors.Wrap(err, "fail to initialize ERC20")
	}

	return &Token{
		ERC20:    token,
		ChainId:  chainId,
		Address:  address,
		Metadata: nil,
		Client:   c,
	}, nil
}

func MustNew(ChainId int, tokenAddress string, c *client.Client) *Token {
	a, err := New(ChainId, tokenAddress, c)
	if err != nil {
		panic(err)
	}
	return a
}
