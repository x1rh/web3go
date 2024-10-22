package client

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetSepoliaSuccessTransactionReceipt(t *testing.T) {
	c := testSuit()

	successTx := "0xb9e94f4c05ee809fa0fbd567842169d309af98678b55b097b66000e153edbba1"
	receipt, err := c.Client.TransactionReceipt(context.Background(), common.HexToHash(successTx))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("receipt: %+v\n", receipt)
}

func TestGetSepoliaFailedTransactionReceipt(t *testing.T) {
	c := testSuit()
	sepoliaFailedTx := "0x633a6b6d96d4414309be65c945a4053bf940b9e4ee964d745bc05a435e273fb7"
	receipt, err := c.Client.TransactionReceipt(context.Background(), common.HexToHash(sepoliaFailedTx))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("receipt: %+v\n", receipt)
}

