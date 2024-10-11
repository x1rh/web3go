package erc20

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (a *Adapter) Decimals(ctx context.Context) (uint64, error) {
	decimals, err := a.ERC20.Decimals(&bind.CallOpts{
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

func (a *Adapter) Symbol(ctx context.Context) (string, error) {
	symbol, err := a.ERC20.Symbol(&bind.CallOpts{
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

func (a *Adapter) Name(ctx context.Context) (string, error) {
	name, err := a.ERC20.Name(&bind.CallOpts{
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
