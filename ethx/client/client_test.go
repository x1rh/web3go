package client

import (
	"testing"
)

func testSuit() *Client {
	url := "https://eth-sepolia.g.alchemy.com/v2/gOeoBV9mlFL1pWj7qbKEdlB6pXTfNum6"
	chainId := 11155111
	chainName := "eth-sepolia"
	return MustNew(Config{
		URL:       url,
		ChainId:   chainId,
		ChainName: chainName,
	})
}

func TestMustNewClient(t *testing.T) {
	testSuit()
}
