package token

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

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

func (t *Token) TokenMetadata() (*TokenMetadata, error) {
	ctx := context.Background()
	name, err := t.Name(ctx)
	if err != nil {
		return nil, err
	}

	symbol, err := t.Symbol(ctx)
	if err != nil {
		return nil, err
	}

	decimals, err := t.Decimals(ctx)
	if err != nil {
		return nil, err
	}

	return &TokenMetadata{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}
