package client

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type spoofTransport struct {
	base   http.RoundTripper
	origin string
}

func (t *spoofTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Origin", t.origin)
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	)
	return t.base.RoundTrip(req)
}

func dialEthClientCtx(ctx context.Context, rpcURL string, origin string) (*ethclient.Client, error) {
	if origin == "" {
		return ethclient.DialContext(ctx, rpcURL)
	}

	rpcEndpoint, err := url.Parse(rpcURL)
	if err != nil {
		return nil, err
	}

	if rpcEndpoint.Scheme != "http" && rpcEndpoint.Scheme != "https" {
		return ethclient.DialContext(ctx, rpcURL)
	}

	httpClient := &http.Client{
		Transport: &spoofTransport{
			base:   http.DefaultTransport,
			origin: origin,
		},
	}

	rpcClient, err := rpc.DialHTTPWithClient(rpcURL, httpClient)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(rpcClient), nil
}
