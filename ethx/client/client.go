package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Client *ethclient.Client
	Config Config
}

func New(c Config, opts ...Option) (*Client, error) {
	return NewCtx(context.Background(), c, opts...)
}

func MustNew(c Config, opts ...Option) *Client {
	client, err := New(c, opts...)
	if err != nil {
		panic(err)
	}
	return client
}

func NewCtx(ctx context.Context, c Config, opts ...Option) (*Client, error) {
	o := applyOptions(opts)
	client, err := dialEthClientCtx(ctx, c.URL, o.origin)
	if err != nil {
		return nil, fmt.Errorf("fail to dial blockchain node: %w", err)
	}
	return &Client{
		Client: client,
		Config: c,
	}, nil
}

func MustNewCtx(ctx context.Context, c Config, opts ...Option) *Client {
	client, err := NewCtx(ctx, c, opts...)
	if err != nil {
		panic(err)
	}
	return client
}
