package client

import (
	"context"
	"testing"
)

func TestClient_Transfer(t *testing.T) { 
	pk := ""  // fill it when test 
	
	if pk == "" {
		return 
	}
	
	to := "0xE3a463d743F762D538031BAD3f1E748BB41f96ec"
	amount := "0.0001"
	cli := testSuit()
	chainId := int64(11155111)
	gasLimit := uint64(21000) 
	
	gasPrice, err := cli.GasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	tx, err := cli.Transfer(context.Background(), chainId, pk, to, amount, gasLimit, gasPrice)
	if err != nil {
		t.Fatal(err)
	} 
	t.Logf("tx detail: %+v\n", tx)	
}
