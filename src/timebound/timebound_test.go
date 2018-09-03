package timebound

import (
	"account"
	"log"
	"testing"
	"time"
)

func TestTxWithTimeBound(t *testing.T) {
	a := account.MakeAccount(sourceAccount)
	tx, err := TxWithTimeBound(uint64(time.Now().Add(time.Second*10).Unix()), 0, a)
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	txHash, err := a.SignAndSubmit(tx)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(txHash)
}
