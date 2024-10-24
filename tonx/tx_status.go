package tonx

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

func DefaultMainnetGetTransaction(lt uint64, address, txHash string) (*tlb.Transaction, error) {
	mainnet := "https://ton.org/global.config.json"
	return GetTransaction(mainnet, lt, address, txHash)
}

func DefaultTestnetGetTransaction(lt uint64, address, txHash string) (*tlb.Transaction, error) {
	testnet := "https://ton-blockchain.github.io/testnet-global.config.json"
	return GetTransaction(testnet, lt, address, txHash)
}

func GetTransaction(url string, lt uint64, walletAddress, txHash string) (*tlb.Transaction, error) {
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), url)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get config")
	}

	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "connection error")
	}

	// initialize ton api lite connection wrapper with full proof checks
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	wa, err := address.ParseAddr(walletAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("fail to parse address: %s", walletAddress))
	}

	txb, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, errors.Wrap(err, "fail to decode tx hash to []byte")
	}

	// list 5 transaction to find the specific tx
	transactionList, err := api.ListTransactions(context.Background(), wa, 5, lt, txb)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("fail to list transactions, lt=%d, address=%s, tx=%s", lt, walletAddress, txHash))
	}

	for _, tx := range transactionList {
		hexTxHash := hex.EncodeToString(tx.Hash)
		fmt.Printf("tx hash: %s\n", hexTxHash)
		if strings.EqualFold(hexTxHash, txHash) {
			return tx, nil
		}
	}
	return nil, fmt.Errorf("fail to find the tx: %s", txHash)
}

// check receiver's address
func CheckTxIn(tx *tlb.Transaction, sender, receiver, amount string) (bool, error) {
	switch tx.Description.Description.(type) {
	default:
		return false, fmt.Errorf("invalid tx description")
	case tlb.TransactionDescriptionOrdinary:
		if tx.IO.In != nil {
			if tx.IO.In.MsgType == tlb.MsgTypeInternal {
				in := tx.IO.In.AsInternal().Amount.Nano()
				intTx := tx.IO.In.AsInternal()
				tonAmount := tlb.FromNanoTON(in).String()
				fromAddress := intTx.SrcAddr

				if tonAmount != amount {
					return false, fmt.Errorf("amount not equal tonAmount: %s != %s", amount, tonAmount)
				}

				senderAddress, err := address.ParseAddr(sender)
				if err != nil {
					return false, errors.Wrap(err, "fail to parse address")
				}

				if !senderAddress.Equals(fromAddress) {
					return false, fmt.Errorf("sender address not equal fromAddress: %s != %s", sender, fromAddress)
				}
				// comment := intTx.Comment()
				// if comment != "" {
				// 	// todo: check comment ?
				// }
			} else if tx.IO.In.MsgType == tlb.MsgTypeExternalIn {
				return false, errors.New("fail to handle tx.IO.In.MsgType")
			}
		}
	}
	return true, nil
}

// check sender's address
func CheckTxOut(tx *tlb.Transaction, sender, receiver, amount string) (bool, error) {
	if tx.IO.Out != nil {
		listOut, err := tx.IO.Out.ToSlice()
		if err != nil {
			return false, errors.Wrap(err, "OUT MESSAGES NOT PARSED DUE TO ERR")
		}

		// NOTICE: I'm assuming that there is only **ONE** outgoing transaction, and it is to the receiver.
		for _, m := range listOut {
			if m.MsgType == tlb.MsgTypeInternal {
				receiverAddressInTx := m.Msg.DestAddr()
				tonAmount := tlb.FromNanoTON(m.AsInternal().Amount.Nano()).String()

				if tonAmount != amount {
					return false, fmt.Errorf("amount not equal tonAmount: %s != %s", amount, tonAmount)
				}

				receiverAddress, err := address.ParseAddr(receiver)
				if err != nil {
					return false, errors.Wrap(err, "fail to parse address")
				}

				if !receiverAddress.Equals(receiverAddressInTx) {
					return false, fmt.Errorf("receiver address not equal tx receiver address: %s != %s", receiverAddress, receiverAddressInTx)
				}
				return true, nil
			}
		}
	}
	return false, errors.New("not found")
}
