package client

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/x1rh/ethx/convertx"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

// Transfer ether to another account
// PrivateKey:
// ToAddress:
// amount: transfer volume
func (c *Client) Transfer(
	ctx context.Context,
	chainID int64,
	PrivateKey string,
	ToAddress string,
	amount string,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	value, err := convertx.EtherToWei(amount)
	if err != nil {
		return nil, errors.Wrap(err, "invalid amount")
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "crypto.HexToECDSA() error")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.Wrap(err, "fail to get public key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.Client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get nonce")
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(ToAddress), value.BigInt(), gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "fail to sign a tx")
	}

	err = c.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to send transaction")
	}

	return signedTx, nil
}
