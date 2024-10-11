package hub

import (
	"github.com/x1rh/ethx/client"
	"github.com/x1rh/ethx/config"
)

type Hub struct {
	chains  map[int]config.Config
	mapping map[int]*client.Adapter
}

func New(chains map[int]config.Config) (*Hub, error) {
	h := &Hub{
		chains:  make(map[int]config.Config),
		mapping: make(map[int]*client.Adapter),
	}

	for k, v := range chains {
		h.chains[k] = v
		c := client.MustNewAdapter(v)
		h.mapping[k] = c
	}

	return h, nil
}

func (h *Hub) WithChainID(chainID int) *client.Adapter{
	return h.mapping[chainID]
}
