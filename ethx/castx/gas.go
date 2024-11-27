package castx

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/x1rh/web3go/ethx/types"
)

// GasFee in wei
func GasFee(gasLimit, gasPrice any) (*decimal.Decimal, error) {
	gasLimitD, err := ToDecimal(gasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "invalid gas limit")
	}

	gasPriceD, err := ToDecimal(gasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "invalid gas price")
	}

	gasFee := gasLimitD.Mul(*gasPriceD)

	return &gasFee, nil
}

func GasFeeInWei(gasLimit, gasPrice any) (*decimal.Decimal, error) {
	return GasFee(gasLimit, gasPrice)
}

func GasFeeInGwei(gasLimit, gasPrice any) (*decimal.Decimal, error) {
	gasFee, err := GasFee(gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}

	return DivByDecimal(gasFee, 9)
}

func GasFeeInEth(gasLimit, gasPrice any) (*decimal.Decimal, error) {
	gasFee, err := GasFee(gasLimit, gasPrice)
	if err != nil {
		return nil, err
	}

	return DivByDecimal(gasFee, 18)
}

func PrettyGasInfo(gasLimit, gasPrice any) (*types.GasInfo, error) {
	_gasLimit, err := ToDecimal(gasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "invalid gas limit")
	}

	_gasPrice, err := ToDecimal(gasPrice)
	if err != nil {
		return nil, errors.Wrap(err, "invalid gas price")
	}

	gasFee, err := GasFee(_gasLimit, _gasPrice)
	if err != nil {
		return nil, err
	}

	gasPriceInGwei, err := DivByDecimal(_gasPrice, 9)
	if err != nil {
		return nil, err
	}

	gasFeeInEth, err := DivByDecimal(gasFee, 18)
	if err != nil {
		return nil, err
	}

	return &types.GasInfo{
		GasLimit:       _gasLimit.BigInt().Uint64(),
		GasPrice:       _gasPrice.BigInt(),
		GasPriceInGwei: gasPriceInGwei,
		GasFeeInEth:    gasFeeInEth,
	}, nil
}
