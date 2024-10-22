package constant

import (
	"math/big"

	"github.com/shopspring/decimal"
)

const (
	GWeiI64 int64 = 1000000000
	EthI64  int64 = 1000000000000000000
)

var (
	GWeiDecimal = decimal.NewFromInt(GWeiI64)
	EthDecimal  = decimal.NewFromInt(EthI64)
)

var (
	// deprecated
	ZeroBI = big.NewInt(0)
	OneBI  = big.NewInt(1)
	ZeroD  = decimal.NewFromInt(0)
	OneD   = decimal.NewFromInt(1)

	BigIntZero  = big.NewInt(0)
	BigIntOne   = big.NewInt(1)
	DecimalZero = decimal.NewFromInt(0)
	DecimalOne  = decimal.NewFromInt(1)
)
