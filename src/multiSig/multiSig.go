package multiSig

import (
	"account"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"log"
)

/**
构建这样一个事务：
从账户A向账户Z发起一笔转账，金额随意
该账户A需要A和B的签名，A和B的权重相等都为1，事务的中级阈值设置为2（转账属于中级操作）
分别测试只A签名、只B签名、A和B签名、A,B和C签名时，事务是否成功
*/

// accounts' seed 已在测试网络激活
// GBUMTL7FJP5QOAQFJY4TBDJPK7ZM63SHJRNSTSL6TR6OW2VL33342Z57
const ASeed string = "SBBIL7ZF2YAZPUBX3GA5CIG7LFT5L5VE3L6I52VMDGLGFXN6AU57AS7Z"

// GB4NWUQS2XZAFEZL3YJMZALE77OCPB6H65OJWBVDYWXM7ZDU26AF5QIT
const BSeed string = "SDG2AGLISYYKJEKWGJQGY7H3FIST26KPLCPVADOCVKFK4RMXLOS2SWTB"

// GAM5LMSP5PNN3V656DRXFBA4AEGVTWUQ4TZYEXKUUYRWNOEKBQ453ITJ
const CSeed string = "SCC64TNJZKLH55OUPY2DF7KRMT2L6LGNZLHQ6MFGPMF62SHY4ED3KJRS"

// GAFPQF3KNHQUTGMXO7HPWBN2YHHUBPQZFD2Q6NA7Y6ZUODB2XE6XOXEX
const ZSeed string = "SAOQSFQGHLHBYHVDUZHBMFDD2U3HQR6Q6H7AGZCNBOGKIK26QEHF34EJ"

// 创建账户A，设置A的Master Key权重位1，设置中级阈值为2，添加B，C为账户签名者
func SetAccountA() {
	aa := account.MakeAccount(ASeed)
	ab := account.MakeAccount(BSeed)
	ac := account.MakeAccount(CSeed)
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: aa.Address()},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.SetOptions(
			build.MasterWeight(1),
			build.SetMediumThreshold(2),      // other level leave 0
			build.AddSigner(ab.Address(), 1), // one SetOption operation only allow one AddSigner parameter
		),
		build.SetOptions(build.AddSigner(ac.Address(), 1)), // so, add another signer in another SetOption operation
	)
	account.PanicIfError(err)

	txHash, err := aa.SignAndSubmit(tx)
	log.Print(txHash)
}

// 发起转账
// 只A或B签名
func OneSignerPay(source, signer, destination string) error {
	from := account.MakeAccount(source)
	sign := account.MakeAccount(signer)
	to := account.MakeAccount(destination)
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: from.Address()},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.Payment(
			build.NativeAmount{"100"},
			build.Destination{to.Address()},
		),
	)
	account.PanicIfError(err)
	txHash, err := sign.SignAndSubmit(tx)
	log.Println("txid:", txHash)
	account.LogError(err)
	return err
}

// A和B签名 A,B和C签名
func MultiSignerPay(accounts ...string) error {
	from := account.MakeAccount(accounts[0]) // from account
	to := account.MakeAccount(accounts[1])   // destination account
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: from.Address()},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.Payment(
			build.NativeAmount{"100"},
			build.Destination{to.Address()},
		),
	)
	account.PanicIfError(err)
	txHash, err := account.MultiSignAndSubmit(tx, accounts[2:]...)
	log.Println("txid:", txHash)
	account.LogError(err)
	return err
}
