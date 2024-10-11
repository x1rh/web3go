package client

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetSepoliaSuccessTransactionReceipt(t *testing.T) {
	successTx := "0xb9e94f4c05ee809fa0fbd567842169d309af98678b55b097b66000e153edbba1"

	adapter := testSuit()
	res, err := adapter.Client.TransactionReceipt(context.Background(), common.HexToHash(successTx))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Type is: ", res.Type)
	t.Log("PostState is: ", res.PostState)
	t.Log("Status is: ", res.Status)
	t.Log("CumulativeGasUsed is: ", res.CumulativeGasUsed)
	t.Log("Bloom is: ", res.Bloom)
	t.Log("Logs is: ", res.Logs)
	t.Log("TxHash is: ", res.TxHash)
	t.Log("ContractAddress is: ", res.ContractAddress)
	t.Log("EffectiveGasPrice is: ", res.EffectiveGasPrice)
	t.Log("BlobGasUsed is: ", res.BlobGasUsed)
	t.Log("BlobGasPrice is: ", res.BlobGasPrice)
	t.Log("BlockNumber is: ", res.BlockNumber)
	t.Log("TransactionIndex is: ", res.TransactionIndex)
}

func TestGetSepoliaFailedTransactionReceipt(t *testing.T) {
	sepoliaFailedTx := "0x633a6b6d96d4414309be65c945a4053bf940b9e4ee964d745bc05a435e273fb7"
	adapter := testSuit()
	res, err := adapter.Client.TransactionReceipt(context.Background(), common.HexToHash(sepoliaFailedTx))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Type is: ", res.Type)
	t.Log("PostState is: ", res.PostState)
	t.Log("Status is: ", res.Status)
	t.Log("CumulativeGasUsed is: ", res.CumulativeGasUsed)
	t.Log("Bloom is: ", res.Bloom)
	t.Log("Logs is: ", res.Logs)
	t.Log("TxHash is: ", res.TxHash)
	t.Log("ContractAddress is: ", res.ContractAddress)
	t.Log("EffectiveGasPrice is: ", res.EffectiveGasPrice)
	t.Log("BlobGasUsed is: ", res.BlobGasUsed)
	t.Log("BlobGasPrice is: ", res.BlobGasPrice)
	t.Log("BlockNumber is: ", res.BlockNumber)
	t.Log("TransactionIndex is: ", res.TransactionIndex)
}
