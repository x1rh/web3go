package manager

import (
	"github.com/pkg/errors"

	"github.com/x1rh/web3go/ethx/app/erc20/token"
	"github.com/x1rh/web3go/ethx/chain"
	"github.com/x1rh/web3go/ethx/client"
)

type TokenManager struct {
	configs map[int]chain.IConfig
	clients map[int]*client.Client
	tokens  map[int]map[string]*token.Token // (chainId, tokenAddress) => token
}

// Get find a token with chainId and address
func (m *TokenManager) Get(chainId int, address string) (*token.Token, error) {
	tokenMap, found := m.tokens[chainId]
	if !found {
		return nil, errors.New("invalid chain id")
	}
	token, found := tokenMap[address]
	if !found {
		return nil, errors.New("token not found")
	}
	return token, nil
}

// Query will query block chain to get the token and cache it
func (tm *TokenManager) Query(chainId int, address string) (*token.Token, error) {
	t, _ := tm.Get(chainId, address)
	if t != nil {
		return t, nil
	}

	// if not found, we will initialize it
	client, found := tm.clients[chainId]
	if !found {
		return nil, errors.New("nil client")
	}

	t, err := token.New(chainId, address, client)
	if err != nil {
		return nil, errors.Wrap(err, "fail to initialize token")
	}

	// NOTICE:
	tm.tokens[chainId][address] = t

	return t, nil
}

// user should gurantee chain config is unique
func New(configs map[int]chain.IConfig) (TokenManager, error) {
	tokenManager := &TokenManager{
		configs: configs,
		clients: make(map[int]*client.Client, len(configs)),
		tokens:  make(map[int]map[string]*token.Token, len(configs)),
	}

	return *tokenManager, nil
}
