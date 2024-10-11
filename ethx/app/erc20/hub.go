package erc20

import (
	"fmt"

	"github.com/x1rh/ethx/chain"
	"github.com/x1rh/ethx/client"
)

type Hub struct {
	chains       map[int64]chain.Config
	EthClientMap map[int64]*client.Client
	ERC20Map     map[int64]map[string]*Adapter
}

func NewHub(chains map[int64]chain.Config) (*Hub, error) {
	p := &Hub{
		chains:       make(map[int64]chain.Config),
		EthClientMap: make(map[int64]*client.Client),
		ERC20Map:     make(map[int64]map[string]*Adapter),
	}

	for chainId, chainConfig := range chains {
		p.chains[chainId] = chains[chainId]
		c := client.MustNewClient(chainConfig)
		p.EthClientMap[chainId] = c
		p.ERC20Map[chainId] = make(map[string]*Adapter)
	}

	return p, nil
}

func (p *Hub) With(chainID int64, tokenAddress string) *Adapter {
	c, ok := p.EthClientMap[chainID]
	if !ok {
		panic(fmt.Sprintf("missing eth client, chainID=%d", chainID))
	}

	erc20Map, ok := p.ERC20Map[chainID]
	if !ok {
		panic(fmt.Sprintf("missing wrapper mapping, chainID=%d", chainID))
	}

	adapter, ok := erc20Map[tokenAddress]
	if !ok {
		a := MustNewAdapter(chainID, tokenAddress, c)
		erc20Map[tokenAddress] = a
		return a
	}
	return adapter
}
