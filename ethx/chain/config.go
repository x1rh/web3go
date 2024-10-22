package chain

type Config struct {
	ChainId   int
	ChainName string
	URL       string
}

func (c *Config) GetChainId() int {
	return c.ChainId
}

func (c *Config) GetChainName() string {
	return c.ChainName
}

func (c *Config) GetURL() string {
	return c.URL
}

var _ IConfig = (*Config)(nil)
