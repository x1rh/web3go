package client

import "strings"

const (
	ClientTypeHTTPS   = "https"
	ClientTypeWSS     = "wss"
	ClientTypeUnknown = "unknown"
)

type Config struct {
	ChainId   int    `json:"chainId"`
	ChainName string `json:"chainName"`
	URL       string `json:"url"`
	Type      string // wss or http
}

func NewConfig(chainId int, chainName string, url string) Config {
	c := Config{
		ChainId:   chainId,
		ChainName: chainName,
		URL:       url,
	}
	if strings.HasPrefix(url, "https") {
		c.Type = ClientTypeHTTPS
	} else if strings.HasPrefix(url, "wss") {
		c.Type = ClientTypeWSS
	} else {
		c.Type = ClientTypeUnknown
	}
	return c
}
