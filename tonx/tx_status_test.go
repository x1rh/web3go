package tonx

import (
	"testing"

	"github.com/xssnick/tonutils-go/address"
)

func TestDefaultTestnetGetTransaction(t *testing.T) {
	// address receive 2 ton
	lt := uint64(26725328000003)
	address := "UQCwxyBrDh2FL3MycXF0-XgInMp1aMbND_1SIggGS2cVjAlV"
	txHash := "746AA9D08540713EEC53646655ED216F5CACCC6CDA042D6BD865FE6B8A2C4C1C"

	tx, err := DefaultTestnetGetTransaction(lt, address, txHash)
	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v\n", tx)
}

func TestDefaultTestnetGetTransaction2(t *testing.T) {
	// address `0QAsR34m_SDd4fbk-vJCwUWNK17AVuRf152vtbXV9C9Tizlf` send 1 ton to `0QCwxyBrDh2FL3MycXF0-XgInMp1aMbND_1SIggGS2cVjLLf`
	lt := uint64(26926071000001)
	senderAddress := "0QAsR34m_SDd4fbk-vJCwUWNK17AVuRf152vtbXV9C9Tizlf"   // tt wallet
	receiverAddress := "0QCwxyBrDh2FL3MycXF0-XgInMp1aMbND_1SIggGS2cVjLLf" // my wallet
	txHash := "0c99cffbc8e237b5e5665f1770e2414c88e1b74a60c8fc464f0cbb69c6e29bfa"

	tx, err := DefaultTestnetGetTransaction(lt, receiverAddress, txHash)
	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v\n", tx)

	amount := "1"
	ok, err := CheckTxIn(tx, senderAddress, receiverAddress, amount)
	if ok {
		t.Logf("check tx success")
	} else {
		t.Error("err", err)
	}
}

func TestDefaultTestnetGetTransaction3(t *testing.T) {
	lt := uint64(27073348000001)
	senderAddress := "0QAsR34m_SDd4fbk-vJCwUWNK17AVuRf152vtbXV9C9Tizlf"   // tt wallet
	receiverAddress := "0QC9EbkIoArFmivwrDf1_se0VelUqAxhxTpQnihzpJvr1zP3" // s wallet
	txHash := "d42b73a782b544253311228996ec7d81badd371be548d68a289aa8a6ed50a4db"

	// 查询发送者的交易hash, 并检查接收者和ton数额
	tx, err := DefaultTestnetGetTransaction(lt, senderAddress, txHash)
	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v\n", tx)

	amount := "0.1"
	ok, err := CheckTxOut(tx, senderAddress, receiverAddress, amount)
	if ok {
		t.Logf("check tx success")
	} else {
		t.Error("err", err)
	}
}

func TestDefaultTestnetGetTransaction4(t *testing.T) {
	lt := uint64(27167288000001)
	senderAddress := "0QAsR34m_SDd4fbk-vJCwUWNK17AVuRf152vtbXV9C9Tizlf"   // tt wallet
	receiverAddress := "0QC9EbkIoArFmivwrDf1_se0VelUqAxhxTpQnihzpJvr1zP3" // s wallet
	txHash := "1266a7ea48b7991de4a0cb1e35447bd07221a78f992a5d91cd795ace64f447b0"

	// 查询发送者的交易hash, 并检查接收者和ton数额
	tx, err := DefaultTestnetGetTransaction(lt, senderAddress, txHash)
	if err != nil {
		t.Error(err)
	}

	t.Logf("tx: %+v\n", tx)

	amount := "0.1"
	ok, err := CheckTxOut(tx, senderAddress, receiverAddress, amount)
	if ok {
		t.Logf("check tx success")
	} else {
		t.Error("err", err)
	}
}

func TestDefaultMainnetGetTransaction(t *testing.T) {

}

func TestWalletAddress(t *testing.T) {
	a1, err := address.ParseRawAddr("0:2c477e26fd20dde1f6e4faf242c1458d2b5ec056e45fd79dafb5b5d5f42f538b")
	if err != nil {
		t.Logf("parse raw address error: %v", err)
	} else {
		t.Logf("address: %+v\n", a1)
	}

	a2, err := address.ParseAddr("0:2c477e26fd20dde1f6e4faf242c1458d2b5ec056e45fd79dafb5b5d5f42f538b")
	if err != nil {
		t.Logf("parse address error: %v", err)
	} else {
		t.Logf("address: %+v\n", a2)
	}
}
