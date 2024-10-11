package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"reflect"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// ParsePrivateKey return privateKey, publicKey, address, error
func ParsePrivateKey(pk string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, *common.Address, error) {
	privateKeyBytes, err := hex.DecodeString(pk)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "fail to decode hex string ")
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "fail to get private key")
	}

	cryptoPublicKey := privateKey.Public()
	publicKey, ok := cryptoPublicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, nil, errors.New("fail to get public key")
	}

	address := crypto.PubkeyToAddress(*publicKey)
	return privateKey, publicKey, &address, nil
}

func GetAddress(privateKey string) (*common.Address, error) {
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "fail to decode private key")
	}

	ethPrivateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get private key")
	}

	publicKey := ethPrivateKey.Public()
	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("fail to get public key")
	}

	address := crypto.PubkeyToAddress(*ecdsaPublicKey)
	return &address, nil
}

var (
	addressRe = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
)

func SimpleCheck(address string) bool {
	return addressRe.MatchString(address)
}

func MustAddress(iaddress any) (*common.Address, error) {
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


func IsZeroAddress(iaddress interface{}) bool {
	addr, err := MustAddress(iaddress)
	if err != nil {
		return false 
	}
	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := addr.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}