package types

import (
	"github.com/shopspring/decimal"
	"math/big"
)

type GasInfo struct {
	GasLimit       uint64
	GasPrice       *big.Int
	GasPriceInGwei *decimal.Decimal
	GasFeeInEth    *decimal.Decimal
}
