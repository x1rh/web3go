package castx

import (
	"testing"
)

func TestToAddress(t *testing.T) {
	address1 := "0x13CB6AE34A13a0977F4d7101eBc24B87Bb23F0d5"
	address2, err := ToAddress(address1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(address2)
}
