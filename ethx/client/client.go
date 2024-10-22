package client

import (
	"github.com/x1rh/web3go/ethx/chain"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Client struct {
	Client *ethclient.Client
	Config chain.IConfig
}

func NewClient(c chain.IConfig) (*Client, error) {
	cli, err := ethclient.Dial(c.GetURL())
	if err != nil {
		return nil, errors.Wrap(err, "fail to dial ethereum node")
	}
	return &Client{
		Client: cli,
		Config: c,
	}, nil
}

func MustNewClient(c chain.IConfig) *Client {
	cli, err := NewClient(c)
	if err != nil {
		panic(err)
	}
	return cli
}
