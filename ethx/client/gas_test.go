package client

import (
	"context"
	"testing"
)


func TestGasPrice(t *testing.T) {
	adapter := testSuit()
	gasPrice, err := adapter.GasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("gas price=%v\n", gasPrice)
}

func TestEstimateGas(t *testing.T) {

}
