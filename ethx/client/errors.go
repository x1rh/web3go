package client

import (
	"errors"
)

var (
	ErrInvalidPublicKey  = errors.New("invalid public key")
	ErrInvalidPrivateKey = errors.New("invalid private key")
	ErrInvalidEthAmount  = errors.New("invalid Ethereum amount")
	ErrSignTx            = errors.New("fail to sign transaction")
	ErrUnknown           = errors.New("unknown error")
)
