package clienthub

import (
	"github.com/pkg/errors"

	"github.com/x1rh/web3go/ethx/chain"
	"github.com/x1rh/web3go/ethx/client"
)

// ClientHub is a multi-chain client manager
type ClientHub struct {
	chains  map[int]chain.IConfig
	clients map[int]*client.Client
}

func New(chains map[int]chain.IConfig) (*ClientHub, error) {
	hub := &ClientHub{
		chains:  make(map[int]chain.IConfig),
		clients: make(map[int]*client.Client),
	}

	for chainId, chainConfig := range chains {
		hub.chains[chainId] = chainConfig
		client, err := client.NewClient(chainConfig)
		if err != nil {
			return nil, errors.Wrap(err, "fail to create client")
		}
		hub.clients[chainId] = client
	}
	return hub, nil
}

func (h *ClientHub) WithChainID(chainId int) (*client.Client, error) {
	client, found := h.clients[chainId]
	if !found {
		return nil, errors.New("invalid chainId")
	}
	return client, nil
}

func (h *ClientHub) MustWithChainID(chainId int) *client.Client {
	client, _ := h.WithChainID(chainId)
	return client
}
