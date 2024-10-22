package chain

type IConfig interface {
	GetChainId() int      // chain id
	GetURL() string       // ethereum node url
	GetChainName() string // chain name
}
