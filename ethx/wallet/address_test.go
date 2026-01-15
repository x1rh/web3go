package wallet

import (
	"testing"
)

func TestGetAddress(t *testing.T) {
	pk := "ab8e0b9ba5a7ae6a1925602dfd50bc26bc0c3fe60673308934f82eb5f18a0b78"
	address := "0x15509Be5DbADF0D039a4DBc0839461D6E90F88A5"

	address2, err := GetAddress(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("address : %s\n", address)
	t.Logf("address2: %s\n", address2)
	if address != address2.Hex() {
		t.Fatal("not equal")
	}
}

func TestGetAddressFuzz(t *testing.T) {
	walletList, err := BatchGenerateWallet(1000)
	if err != nil {
		t.Fatal(err)
	}

	for _, w := range walletList {
		t.Logf("%+v\n", w)
		address, err := GetAddress(w.PrivateKey)
		if err != nil {
			t.Fatal(err)
		}
		if w.AddressHex != address.Hex() {
			t.Fatal("not equal")
		}
	}
}

func TestSimpleCheck(t *testing.T) {
	t.Logf("is valid: %v\n", addressRe.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
}
