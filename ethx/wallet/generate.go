package wallet

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Wallet struct {
	PrivateKey string
	Address    string
}

func GenerateWallet() (*Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.Wrap(err, "fail to generate key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {	
		return nil, errors.New("fail to cast public key to ECDSA") 
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	pk := hexutil.Encode(privateKeyBytes)[2:]
	
	return &Wallet{
		PrivateKey: pk,
		Address:    address,
	}, nil
}

func BatchGenerateWallet(n int) ([]*Wallet, error) {
	var wallets []*Wallet
	for i := 0; i < n; i++ {
		w, err := GenerateWallet()
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}

	return wallets, nil
}
