package constant

import (
	"github.com/shopspring/decimal"
	"math/big"
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
	ZeroBI = big.NewInt(0)
	OneBI  = big.NewInt(1)
	ZeroD  = decimal.NewFromInt(0)
	OneD   = decimal.NewFromInt(1)
)
