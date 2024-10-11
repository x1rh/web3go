package client

import (
	"testing"
	"github.com/x1rh/ethx/config"
)

func testSuit() *Adapter {
	url := "https://eth-sepolia.g.alchemy.com/v2/gOeoBV9mlFL1pWj7qbKEdlB6pXTfNum6"
	chainId := 11155111
	chainName := "eth-sepolia"
	return MustNewAdapter(config.Config{
		URL:       url,
		ChainID:   chainId,
		ChainName: chainName,
	})
}

func TestMustNewClient(t *testing.T) {
	testSuit()
}
