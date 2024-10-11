package convertx

import (
	"github.com/shopspring/decimal"
)

func WeiToGwei(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 9)
}

func WeiToEth(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 18)
}

func GweiToEth(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 9)
}

func GweiToWei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 9)
}

func EthToWei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 18)
}

func EthToGwei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 9)
}
