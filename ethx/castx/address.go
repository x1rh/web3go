package castx

import (
	"github.com/ethereum/go-ethereum/common"
	"errors"
)

func ToAddress(iaddress any) (*common.Address, error) {
	var addr common.Address
	switch v := iaddress.(type) {
	case string:
		addr = common.HexToAddress(v) 
	case common.Address:
		addr = v
	case *common.Address:
		return v, nil 
	default:
		return nil, errors.New("invalid address")
	}
	return &addr, nil 
}

func MustToAddress(iaddress any) *common.Address {
	address, err := ToAddress(iaddress)
	if err != nil {
		panic(err)
	}
	return address
}
