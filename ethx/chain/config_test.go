package chain

import "testing"

func TestConfig(t *testing.T) {
	chainId := 1
	chainName := "ethereum"
	chainURL := "https://eth.drpc.org"
	c := Config{
		ChainId:   chainId,
		ChainName: chainName,
		URL:       chainURL,
	}
	if c.GetChainId() != 1 {
		panic("not equal")
	}
	if c.GetChainName() != "ethereum" {
		panic("not equal")
	}
	if c.GetURL() != chainURL {
		panic("not equal")
	}
}
