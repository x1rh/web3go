package wallet

import (
	"testing"
)

type TestCase struct {
	signature string
	address   string
	message   string
	wanted    bool
}

var cases = []TestCase{
	{
		signature: "0x52bf6c5ff439cca398cdca95c603bedf15e07b594777eb4b6b1dc773467267cc1c11bb75e6a1ae4c7411377881d09261cde1f89c36116396a6eaea9b7c4123a11b",
		address:   "0xb34C0CFAC19819524892E09Afda7402E57CbcDA6",
		message:   "1701432839071",
		wanted:    true,
	},
	//{
	//	signature: "0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301",
	//	address:   "0x96216849c49358B10257cb55b28eA603c874b05E",
	//	message:   "hello",
	//	wanted:    true,
	//},
	{
		signature: "0x0498c6564863c78e663848b963fde1ea1d860d5d882d2abdb707d1e9179ff80630a4a71705da534a562c08cb64a546c6132de26eb77a44f086832cbc1dbe01f71b",
		address:   "0xb052C02346F80cF6ae4DF52c10FABD3e0aD24d81",
		message:   "hello",
		wanted:    true,
	},
	{
		signature: "0xfee7fbb3c69ea19865287c313986a5d7a9c2729b131822594555ad1166dced87069eba4afd3020f835e448e40ee91117923a87158006e0e4f27f9b5b4c0c153d1c",
		message:   "1701843584596-127.0.0.1:39544-fln2dx444x",
		address:   "0xb34C0CFAC19819524892E09Afda7402E57CbcDA6",
		wanted:    true,
	},
	{
		signature: "0xfee7fbb3c69ea19865287c313986a5d7a9c2729b131822594555ad1166dced87069eba4afd3020f835e448e40ee91117923a87158006e0e4f27f9b5b4c0c153d1c",
		message:   "1701843584596-127.0.0.1:39544-fln2dx444x",
		address:   "0xb34C0CFAC19819524892E09Afda7402E57CbcDA6",
		wanted:    true,
	},
	{
		signature: "0xf90b6a5ec633b5b32968d245cf6580fa2365389f2b2d54b4112350a3042eff89307516ad0d520f70d79e265bce7244a6354fe23ccb1a777a1bcde4cb18012eae1b",
		address:   "0x89b6b40b06262414bd3fbf777880b6ebdd26176b",
		message:   "1705913616549--f7cjcfdsvk",
		wanted:    true,
	},
}

func TestVerifySignature(t *testing.T) {
	for _, c := range cases {
		ok, err := VerifySignature(c.address, c.signature, c.message)
		if err != nil {
			t.Fatal(err)
		}
		if ok != c.wanted {
			t.Fatalf("fail %+v\n", c)
		}
	}
}
