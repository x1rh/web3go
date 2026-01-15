package token

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/x1rh/web3go/ethx/client"
)

type TokenInfo struct {
	ChainId  int
	Address  string
	Metadata *TokenMetadata
}

type TokenMetadata struct {
	Name     string
	Symbol   string
	Decimals uint64
}

type Token struct {
	*ERC20

	ChainId int
	Address string

	Client *client.Client

	// we will initialize Metadata after calling `TokenMetadata()` at first time
	// read the data from block chain and cached it
	Metadata *TokenMetadata
}

func New(chainId int, address string, c *client.Client) (*Token, error) {
	token, err := NewERC20(common.HexToAddress(address), c.Client)
	if err != nil {
		return nil, errors.Wrap(err, "fail to initialize ERC20")
	}

	return &Token{
		ERC20:    token,
		ChainId:  chainId,
		Address:  address,
		Metadata: nil,
		Client:   c,
	}, nil
}

func MustNew(ChainId int, tokenAddress string, c *client.Client) *Token {
	a, err := New(ChainId, tokenAddress, c)
	if err != nil {
		panic(err)
	}
	return a
}

func (t *Token) Decimals(ctx context.Context) (uint64, error) {
	decimals, err := t.ERC20.Decimals(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		BlockHash:   common.Hash{},
		Context:     ctx,
	})
	if err != nil {
		return 0, errors.Wrap(err, "fail to get token decimal")
	}
	return uint64(decimals), nil
}

func (t *Token) Symbol(ctx context.Context) (string, error) {
	symbol, err := t.ERC20.Symbol(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		BlockHash:   common.Hash{},
		Context:     ctx,
	})
	if err != nil {
		return "", errors.Wrap(err, "fail to get token symbol")
	}
	return symbol, nil
}

func (t *Token) Name(ctx context.Context) (string, error) {
	name, err := t.ERC20.Name(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		BlockHash:   common.Hash{},
		Context:     ctx,
	})
	if err != nil {
		return "", errors.Wrap(err, "fail to get token name")
	}
	return name, nil
}

func (t *Token) GetMetadata(ctx context.Context) (*TokenMetadata, error) {
	if t.Metadata != nil {
		return &TokenMetadata{
			Name:     t.Metadata.Name,
			Symbol:   t.Metadata.Symbol,
			Decimals: t.Metadata.Decimals,
		}, nil
	}

	name, err := t.Name(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to get name: %w", err)
	}

	symbol, err := t.Symbol(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to get symbol: %w", err)
	}

	decimals, err := t.Decimals(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to get decimals: %w", err)
	}

	return &TokenMetadata{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}

func (t *Token) Approve(
	pk string,
	spender string,
	amount *big.Int,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(t.ChainId)))
	if err != nil {
		return nil, err
	}
	txOpts.GasLimit = gasLimit
	txOpts.GasPrice = gasPrice

	return t.ERC20.Approve(txOpts, common.HexToAddress(spender), amount)
}

func Initialize(ctx context.Context, c *client.Client, ) (*Token, error) {
	return nil, nil
}
