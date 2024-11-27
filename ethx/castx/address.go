package castx

import (
	"github.com/ethereum/go-ethereum/common"
	"errors"
)

func ToAddress(iaddress any) (*common.Address, error) {
	var address common.Address
    switch v := iaddress.(type) {
    case string:
        address = common.HexToAddress(v)
		return &address, nil 
    case common.Address:
		return &v, nil 
    default:
        return nil, errors.New("invalid address type")
    }
}

func MustToAddress(iaddress any) *common.Address {
	address, err := ToAddress(iaddress)
	if err != nil {
		panic(err)
	}
	return address
}
