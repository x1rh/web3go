package ethscan 

type IEthScan interface {
	VerifyContract() 
}

type Config struct {
	URL string   // ethscan url
	Apikey string  // ethscan apikey
}

type EthScan struct {

}