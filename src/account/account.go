package account

import (
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"log"
)

type Account struct {
	seed string
}

func MakeAccount(seed string) (a Account) {
	return Account{seed: seed}
}

func (a *Account) Address() string {
	kp, err := keypair.Parse(a.seed)
	PanicIfError(err)
	return kp.Address()
}
func (a *Account) SignAndSubmit(tx *build.TransactionBuilder) (txHash string, err error) {
	txe, err := tx.Sign(a.seed)
	PanicIfError(err)

	txeB64, _ := txe.Base64()
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	log.Print(resp)
	LogIfErrorMsg(err, GetResultCodeFromError(err))
	return resp.Hash, nil
}

func MultiSignAndSubmit(tx *build.TransactionBuilder, signerSeed ...string) (txHash string, err error) {
	txe, err := tx.Sign(signerSeed...)
	PanicIfError(err)

	txeB64, _ := txe.Base64()
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	log.Print(resp)
	LogIfErrorMsg(err, GetResultCodeFromError(err))
	return resp.Hash, nil
}

func GetResultCodeFromError(err error) string {
	herr, isHorizonError := err.(*horizon.Error)
	if isHorizonError {
		resultCodes, err := herr.ResultCodes()
		if err != nil {
			log.Println("failed to extract result codes from horizon response")
			return ""
		}
		return resultCodes.TransactionCode
	}
	return ""
}

func PanicIfError(e error) {
	PanicErrorMsg(e, "")
}

func PanicErrorMsg(e error, msg string) {
	if e != nil {
		panic(e.Error() + "\n" + msg)
	}
}

func LogError(e error) {
	LogIfErrorMsg(e, "")
}

func LogIfErrorMsg(e error, msg string) {
	if e != nil {
		log.Println(e)
		log.Println(msg)
	}
}
