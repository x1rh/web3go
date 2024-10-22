package tonx

import (
	"context"
	"encoding/hex"
	"fmt"
)

func DefaultMainnetGetTransaction(lt uint64, address, txHash string) {
	mainnet := "https://ton.org/global.config.json"
	return GetTransaction(mainnet, lt, address, txHash)
}

func DefaultTestnetGetTransaction(lt uint64, address, txHash string) {
	testnet := "https://ton-blockchain.github.io/testnet-global.config.json"
	return GetTransaction(testnet, lt, address, txHash)
}

func GetTransaction(url string, lt uint64, address, txHash string) (*tlb.Transaction, error) {
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), url)
	if err != nil {
		return nil, errors.Wrap("fail to get config", err)
	}

	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap("connection error", err)
	}

	// initialize ton api lite connection wrapper with full proof checks
	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)
	transactionList, err := api.ListTransactions(context.Background(), address, 5, lt, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "fail to list transactions")
	}

	for _, tx := range transactionList {
		hexTxHash := hex.EncodeToString(tx.Hash)
		fmt.Printf("tx hash: %s\n", hexTxHash)
		if hexTxHash == txHash {
			return tx, nil
		}
	}
	return nil, errors.Wrap(err, "fail to find the tx")
}
