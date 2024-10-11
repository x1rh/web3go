package convertx

import (
	"github.com/shopspring/decimal"
)

func WeiToGwei(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 9)
}

func WeiToEther(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 18)
}

func GweiToEther(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 9)
}

func GweiToWei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 9)
}

func EtherToWei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 18)
}

func EtherToGwei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 9)
}
