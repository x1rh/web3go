package client

import (
	"github.com/x1rh/ethx/config"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type Adapter struct {
	Client *ethclient.Client
	Config config.Config
}

func NewAdapter(c config.Config) (*Adapter, error) {
	cli, err := ethclient.Dial(c.URL)
	if err != nil {
		return nil, errors.Wrap(err, "dial ethereum error")
	}
	return &Adapter{
		Client: cli,
		Config: c,
	}, nil
}

func MustNewAdapter(c config.Config) *Adapter {
	cli, err := NewAdapter(c)
	if err != nil {
		panic(err)
	}
	return cli
}
