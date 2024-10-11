package wallet

import (
	"testing"
)

func TestBatchGenerateWallet(t *testing.T) {
	numWallets := 5
	wallets, err := BatchGenerateWallet(numWallets)
	if err != nil {
		t.Fatal("Error generating wallets:", err)
	}

	for i, wallet := range wallets {
		t.Logf("Wallet %d:\n", i+1)
		t.Logf("Address: %s\n", wallet.Address)
		t.Logf("Private Key: %s\n", wallet.PrivateKey)
		t.Logf("------------------------------------------------------")
	}
}
