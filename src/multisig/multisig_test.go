package multisig

import "testing"

func TestSetAccountA(t *testing.T) {
	SetAccountA()
}

func TestOneSignerPay(t *testing.T) {
	// all tx should be failed
	err := OneSignerPay(ASeed, ASeed, ZSeed)
	if err == nil {
		t.Error(err)
	}
	err = OneSignerPay(ASeed, BSeed, ZSeed)
	if err == nil {
		t.Error(err)
	}
	err = OneSignerPay(ASeed, CSeed, ZSeed)
	if err == nil {
		t.Error(err)
	}
}

func TestMultiSignerPay(t *testing.T) {
	// A and B and C sign tx
	err := MultiSignerPay(ASeed, ZSeed, ASeed, BSeed, CSeed) // tx_bad_auth_extra error
	if err == nil {
		t.Error(err)
	}
	err = MultiSignerPay(ASeed, ZSeed, ASeed, BSeed) // should be pass
	if err != nil {
		t.Error(err)
	}
}
