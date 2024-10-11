package convertx

import (
	"math/big"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func ToDecimal(x any) (*decimal.Decimal, error) {
	var result decimal.Decimal
	var err error

	switch v := x.(type) {
	case int:
		result = decimal.NewFromInt(int64(v))
	case int64:
		result = decimal.NewFromInt(v)
	case uint64:
		result = decimal.NewFromUint64(v)
	case float32:
		result = decimal.NewFromFloat32(v)
	case float64:
		result = decimal.NewFromFloat(v)
	case *big.Int:
		result = decimal.NewFromBigInt(v, 0)
	case big.Int:
		result = decimal.NewFromBigInt(&v, 0)
	case string:
		result, err = decimal.NewFromString(v)
	case decimal.Decimal:
		result = v
	case *decimal.Decimal:
		return v, nil
	default:
		return nil, errors.New("invalid type")
	}

	return &result, err
}

// MulByDecimal  return x * 10^d
func MulByDecimal(v, d any) (*decimal.Decimal, error) {
	valueD, err := ToDecimal(v)
	if err != nil {
		return nil, errors.Wrap(err, "invalid value")
	}

	decimalD, err := ToDecimal(d)
	if err != nil {
		return nil, errors.Wrap(err, "invalid factor")
	}

	power := decimal.NewFromInt(10).Pow(*decimalD)
	result := valueD.Mul(power)

	return &result, nil
}

// DivByDecimal return x / 10^d
func DivByDecimal(v, d any) (*decimal.Decimal, error) {
	valueD, err := ToDecimal(v)
	if err != nil {
		return nil, errors.Wrap(err, "invalid value")
	}

	decimalD, err := ToDecimal(d)
	if err != nil {
		return nil, errors.Wrap(err, "invalid factor")
	}

	power := decimal.NewFromInt(10).Pow(*decimalD)
	decimal.DivisionPrecision = 64 // notice: trick
	result := valueD.Div(power)

	return &result, nil
}
