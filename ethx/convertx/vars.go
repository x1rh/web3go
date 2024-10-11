package convertx

import "github.com/shopspring/decimal"

var (
	zeroD             = decimal.NewFromInt(0)
	oneD              = decimal.NewFromInt(1)
	f1                = decimal.NewFromFloat(1.23)
	wei               = decimal.NewFromInt(1)
	gwei              = decimal.NewFromInt(1000000000)
	ether             = decimal.NewFromInt(1000000000000000000)
	oneGWeiInEther, _ = decimal.NewFromString("0.000000001")
	oneWeiInEther, _  = decimal.NewFromString("0.000000000000000001")
)

var (
	defaultDivisionPrecision = 64
)

func init() {
	decimal.DivisionPrecision = defaultDivisionPrecision
}
