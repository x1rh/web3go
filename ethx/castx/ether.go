package castx

import (
	"github.com/shopspring/decimal"
)

func WeiToGwei(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 9)
}

func MustWeiToGwei(amount any) *decimal.Decimal {
	d, err := WeiToGwei(amount)
	if err != nil {
		panic(err)
	}
	return d
}

func WeiToEther(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 18)
}

func MustWeiToEther(amount any) *decimal.Decimal {
	d, err := WeiToEther(amount)
	if err != nil {
		panic(err)
	}
	return d 
}

func GweiToEther(amount any) (*decimal.Decimal, error) {
	return DivByDecimal(amount, 9)
}

func MustGweiToEther(amount any) *decimal.Decimal {
	d, err := GweiToEther(amount)
	if err != nil {
		panic(err)
	}
	return d 
}

func GweiToWei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 9)
}

func MustGweiToWei(amount any) *decimal.Decimal {
	d, err := GweiToWei(amount)
	if err != nil {
		panic(err)
	}
	return d 
}

func EtherToWei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 18)
}

func MustEtherToWei(amount any) *decimal.Decimal {
	d, err := EtherToWei(amount)
	if err != nil {
		panic(err)
	}
	return d 
}


func EtherToGwei(amount any) (*decimal.Decimal, error) {
	return MulByDecimal(amount, 9)
}

func MustEtherToGwei(amount any) *decimal.Decimal {
	d, err := EtherToGwei(amount)
	if err != nil {
		panic(err)
	}
	return d 
}


