package tonx

import (
	"context"
	"testing"
)

func TestGeneratePayload(t *testing.T) {
	server, err := NewServer(ServerConfig{
		Secret:          "D.Z & RLX",
		LifeTimePayload: 3000,
		LifeTimeProof:   3000,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	payload, err := server.GeneratePayload()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("server.GeneratePayload() is:%s", payload)

	res, err := server.server.CheckPayload(payload)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("server.GeneratePayload() is:%v", res)
}

func TestVerify(t *testing.T) {
	payload := "e7536e42abbd3ab50000000066def39dc7e4ebc605eec412158e7b77faebd4ce"
	timestamp := int64(1725887424)
	address := "UQCwxyBrDh2FL3MycXF0-XgInMp1aMbND_1SIggGS2cVjAlV"
	domainValue := "t.me"
	signature := "lMs99BV+kuGF3Y7staWk8UUVZC1vvK8Qp5o2yYW6sc+Cyp/Odhb819FiO1i065ve0fPjd3+fMSYSpZ4j/BldDQ=="
	stateInit := "te6cckECFgEAArEAAgE0AgEAUYAAAAA///+IpTOaP32B5DsPSIyuA3gQRDeWpE/fpwDWaH8QgiN1sxqgART/APSkE/S88sgLAwIBIAYEAQLyBQEeINcLH4IQc2lnbrry4Ip/EQIBSBAHAgEgCQgAGb5fD2omhAgKDrkPoCwCASANCgIBSAwLABGyYvtRNDXCgCAAF7Ml+1E0HHXIdcLH4AIBbg8OABmvHfaiaEAQ65DrhY/AABmtznaiaEAg65Drhf/AAtzQINdJwSCRW49jINcLHyCCEGV4dG69IYIQc2ludL2wkl8D4IIQZXh0brqOtIAg1yEB0HTXIfpAMPpE+Cj6RDBYvZFb4O1E0IEBQdch9AWDB/QOb6ExkTDhgEDXIXB/2zzgMSDXSYECgLmRMOBw4hIRAeaO8O2i7fshgwjXIgKDCNcjIIAg1yHTH9Mf0x/tRNDSANMfINMf0//XCgAK+QFAzPkQmiiUXwrbMeHywIffArNQB7Dy0IRRJbry4IVQNrry4Ib4I7vy0IgikvgA3gGkf8jKAMsfAc8Wye1UIJL4D95w2zzYEgP27aLt+wL0BCFukmwhjkwCIdc5MHCUIccAs44tAdcoIHYeQ2wg10nACPLgkyDXSsAC8uCTINcdBscSwgBSMLDy0InXTNc5MAGk6GwShAe78uCT10rAAPLgk+1V4tIAAcAAkVvg69csCBQgkXCWAdcsCBwS4lIQseMPINdKFRQTABCTW9sx4ddM0AByMNcsCCSOLSHy4JLSAO1E0NIAURO68tCPVFAwkTGcAYEBQNch1woA8uCO4sjKAFjPFsntVJPywI3iAJYB+kAB+kT4KPpEMFi68uCR7UTQgQFB1xj0BQSdf8jKAEAEgwf0U/Lgi44UA4MH9Fvy4Iwi1woAIW4Bs7Dy0JDiyFADzxYS9ADJ7VQTG+SC"

	server, err := NewServer(ServerConfig{
		Secret:          "klafdsldssskdsssdd",
		LifeTimePayload: 30000000000000000,
		LifeTimeProof:   30000000000000000,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := server.server.CheckPayload(payload)
	if err != nil {
		t.Fatal(err)
		return
	}

	proof := Params{
		Timestamp:         timestamp,
		DomainLengthBytes: int64(len(domainValue)),
		Address:           address,
		DomainValue:       domainValue,
		Signature:         signature,
		Payload:           payload,
		StateInit:         stateInit,
	}
	t.Logf("%+v\n", proof)
	isVerify, err := server.VerifySignature(context.Background(), &proof)
	if err != nil {
		t.Fatal(err)
	}
	if !isVerify {
		message := "invalid signature"
		t.Fatal(message)
	}

	t.Logf("server.GeneratePayload() is:%v", res)
}

func TestGetNonBounceableAddressByHex(t *testing.T) {
	hex := "0:2c477e26fd20dde1f6e4faf242c1458d2b5ec056e45fd79dafb5b5d5f42f538b"

	Uaddress, err := new(Server).GetNonBounceableAddressByHex(hex)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("Uaddress: ", Uaddress)
	t.Log("isTrue: ", Uaddress == "UQAsR34m_SDd4fbk-vJCwUWNK17AVuRf152vtbXV9C9Ti4LV")
}
