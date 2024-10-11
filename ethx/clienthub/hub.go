package hub

import (
	"github.com/pkg/errors"

	"github.com/x1rh/web3go/ethx/chain"
	"github.com/x1rh/web3go/ethx/client"
)

type ClientHub struct {
	chains  map[int]chain.Config
	clients map[int]*client.Client
}

func New(chains map[int]chain.Config) (*ClientHub, error) {
	h := &ClientHub{
		chains:  make(map[int]chain.Config),
		clients: make(map[int]*client.Client),
	}

	for k, v := range chains {
		h.chains[k] = v
		c, err := client.NewClient(v)
		if err != nil {
			return nil, errors.Wrap(err, "fail to new client")
		}
		h.clients[k] = c
	}
	return h, nil
}

func (h *ClientHub) WithChainID(chainId int) (*client.Client, error) {
	c, found := h.clients[chainId]
	if !found {
		return nil, errors.New("invalid chainId")
	}
	return c, nil
}

func (h *ClientHub) MustWithChainID(chainId int) *client.Client {
	c, _ := h.WithChainID(chainId)
	return c
}
