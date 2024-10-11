package wallet

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// VerifySignature verify signature's ethAddr
func VerifySignature(walletAddress, signature, message string) (bool, error) {
	// convert hex string to byte slice
	sig := hexutil.MustDecode(signature)

	////Transform yellow paper V from 27/28 to 0/1
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27
	}

	msg := accounts.TextHash([]byte(message))
	ecdsaPublicKey, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false, errors.Wrap(err, "fail to parse signature to public key")
	}


	lowerEthAddr := strings.ToLower(walletAddress)    // notice
	recoveredAddr := crypto.PubkeyToAddress(*ecdsaPublicKey)
	lowerRecoveredAddr := strings.ToLower(recoveredAddr.Hex())

	return lowerEthAddr == lowerRecoveredAddr, nil 
}
