package client

import (
	"context"
	"testing"
)

func TestBalance(t *testing.T) {
	addressList := []string{
		"0x53e08d17Ac31e3F152A97F57018D9BB1d4C69cE0",
		"0xE3a463d743F762D538031BAD3f1E748BB41f96ec",
	}

	cli := testSuit()
	for _, address := range addressList {
		balance, err := cli.Balance(context.Background(), address)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s: %s\n", address, balance.String())
	}
}
