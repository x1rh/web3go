package client

import (
	"context"
	"testing"
)

func TestGasPrice(t *testing.T) {
	cli := testSuit()
	gasPrice, err := cli.GasPrice(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("gas price=%v\n", gasPrice)
}

func TestEstimateGas(t *testing.T) {

}
