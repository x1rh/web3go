package client

import (
	"github.com/x1rh/web3go/ethx/chain"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Client struct {
	Client *ethclient.Client
	Config chain.Config
}

func NewClient(c chain.Config) (*Client, error) {
	cli, err := ethclient.Dial(c.URL)
	if err != nil {
		return nil, errors.Wrap(err, "dial ethereum error")
	}
	return &Client{
		Client: cli,
		Config: c,
	}, nil
}

func MustNewClient(c chain.Config) *Client {
	cli, err := NewClient(c)
	if err != nil {
		panic(err)
	}
	return cli
}
