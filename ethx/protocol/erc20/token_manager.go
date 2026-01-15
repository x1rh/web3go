package token

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/x1rh/web3go/ethx/client"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

type TokenManager struct {
	ChainId   int
	client    *client.Client
	tokenMap  map[string]*Token // (chainId, tokenAddress) => token
}

// Get find a token with chainId and address
func (m *TokenManager) Get(address string) (*Token, error) {
	token, found := m.tokenMap[address]
	if !found {
		return nil, ErrTokenNotFound
	}
	return token, nil
}

// Query will query block chain to get the token and cache it
func (tm *TokenManager) Query(ctx context.Context, address string) (*Token, error) {
	t, _ := tm.Get(address)
	if t != nil {
		return t, nil
	}

	// if token not found, we will initialize it
	t, err := New(tm.ChainId, address, tm.client)
	if err != nil {
		return nil, fmt.Errorf("fail to initialize token: %w", err)
	}
	tm.tokenMap[address] = t

	return t, nil
}

// user should gurantee chain config is unique
func NewTokenManager(chainId int, chainName string, client *client.Client) (TokenManager, error) {
	tokenManager := &TokenManager{
		ChainId:   chainId,
		client:    client,
		tokenMap:  make(map[string]*Token),
	}
	return *tokenManager, nil
}
