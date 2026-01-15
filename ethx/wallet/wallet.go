package wallet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/x1rh/web3go/ethx/castx"
)

const (
	defaultReceiptTimeoutSeconds = 300
	receiptRetryAttempts         = 3
	receiptRetryDelay            = time.Second
)

type GasConfig struct {
	GasLimit      uint64
	GasPriceGwei  string
	GasFeeCapGwei string
	GasTipCapGwei string
}

type Wallet struct {
	PrivateKey string
	AddressHex string

	client         *ethclient.Client
	chainID        *big.Int
	from           common.Address
	key            *ecdsa.PrivateKey
	gas            GasConfig
	rpcTimeout     time.Duration
	receiptTimeout time.Duration // timeout for waiting tx to be mined
	mu             sync.Mutex    // protects nonce and serializes tx execution
	nextNonce      *uint64
}

func NewWallet(
	client *ethclient.Client,
	chainID *big.Int,
	key *ecdsa.PrivateKey,
	gas GasConfig,
	rpcTimeoutSeconds int,
	receiptTimeoutSeconds int,
) *Wallet {
	if receiptTimeoutSeconds <= 0 {
		receiptTimeoutSeconds = defaultReceiptTimeoutSeconds
	}
	from := crypto.PubkeyToAddress(key.PublicKey)
	return &Wallet{
		client:         client,
		chainID:        chainID,
		from:           from,
		key:            key,
		gas:            gas,
		rpcTimeout:     time.Duration(rpcTimeoutSeconds) * time.Second,
		receiptTimeout: time.Duration(receiptTimeoutSeconds) * time.Second,
		AddressHex:     from.Hex(),
	}
}

func (w *Wallet) Address() common.Address {
	return w.from
}

func (w *Wallet) WaitReceipt(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	var lastErr error
	for attempt := 0; attempt < receiptRetryAttempts; attempt++ {
		waitCtx, cancel := context.WithTimeout(ctx, w.receiptTimeout)
		receipt, waitErr := bind.WaitMined(waitCtx, w.client, tx)
		cancel()
		if waitErr == nil {
			return receipt, nil
		}
		if waitCtx.Err() == context.DeadlineExceeded {
			lastErr = fmt.Errorf("wait mined timeout after %v", w.receiptTimeout)
		} else {
			lastErr = fmt.Errorf("wait mined: %w", waitErr)
		}
		if attempt < receiptRetryAttempts-1 {
			select {
			case <-time.After(receiptRetryDelay):
			case <-ctx.Done():
				return nil, fmt.Errorf("wait receipt canceled: tx_hash=%s: %w", tx.Hash().Hex(), ctx.Err())
			}
		}
	}
	return nil, fmt.Errorf("wait receipt failed after retries: tx_hash=%s: %w", tx.Hash().Hex(), lastErr)
}

// Lock acquires the tx lock to serialize execution.
// The caller must call Unlock after the tx finishes.
func (w *Wallet) Lock() {
	w.mu.Lock()
}

// Unlock releases the tx lock.
func (w *Wallet) Unlock() {
	w.mu.Unlock()
}

// NewTransactor requires the caller to hold the lock; it does not lock internally.
func (w *Wallet) NewTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := w.nextTxNonce(ctx)
	if err != nil {
		return nil, fmt.Errorf("get nonce: %w", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(w.key, w.chainID)
	if err != nil {
		return nil, fmt.Errorf("new transactor: %w", err)
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	w.applyGasSettings(opts)
	return opts, nil
}

// ResetNonce clears the cached nonce; the next call fetches from chain.
// The caller should already hold the lock.
func (w *Wallet) ResetNonce() {
	w.nextNonce = nil
}

func (w *Wallet) nextTxNonce(ctx context.Context) (uint64, error) {
	if w.nextNonce == nil {
		timeoutCtx, cancel := context.WithTimeout(ctx, w.rpcTimeout)
		defer cancel()

		nonce, err := w.client.PendingNonceAt(timeoutCtx, w.from)
		if err != nil {
			return 0, fmt.Errorf("pending nonce at: %w", err)
		}
		w.nextNonce = &nonce
	}
	nonce := *w.nextNonce
	*w.nextNonce++
	return nonce, nil
}

func (w *Wallet) applyGasSettings(opts *bind.TransactOpts) {
	if w.gas.GasLimit > 0 {
		opts.GasLimit = w.gas.GasLimit
	}
	if w.gas.GasPriceGwei != "" {
		if value, err := castx.GweiToWei(w.gas.GasPriceGwei); err == nil {
			opts.GasPrice = value.BigInt()
		}
		return
	}
	if w.gas.GasFeeCapGwei != "" {
		if value, err := castx.GweiToWei(w.gas.GasFeeCapGwei); err == nil {
			opts.GasFeeCap = value.BigInt()
		}
	}
	if w.gas.GasTipCapGwei != "" {
		if value, err := castx.GweiToWei(w.gas.GasTipCapGwei); err == nil {
			opts.GasTipCap = value.BigInt()
		}
	}
}

// Balance get current balance in wei
func (w *Wallet) Balance(ctx context.Context) (*big.Int, error) {
	balance, err := w.client.BalanceAt(ctx, w.from, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return balance, nil
}

// EtherBalance get current balance in eth
func (w *Wallet) EthBalance(ctx context.Context) (*decimal.Decimal, error) {
	balanceInWei, err := w.Balance(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get balance error")
	}
	return castx.DivByDecimal(balanceInWei, 18)
}

// Transfer sends native token to another account.
// ToAddress:
// amount: transfer volume
func (w *Wallet) Transfer(
	ctx context.Context,
	toAddress string,
	amount string,
	gasLimit uint64,
	gasPrice *big.Int,
) (*types.Transaction, error) {
	value, err := castx.EtherToWei(amount)
	if err != nil {
		return nil, errors.Wrap(err, "invalid amount")
	}

	nonce, err := w.client.PendingNonceAt(ctx, w.from)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get nonce")
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value.BigInt(), gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(w.chainID), w.key)
	if err != nil {
		return nil, errors.Wrap(err, "fail to sign a tx")
	}

	err = w.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to send transaction")
	}

	return signedTx, nil
}
