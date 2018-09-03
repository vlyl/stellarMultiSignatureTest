package timebound

import (
	"account"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// GC53V73G5P7NDANSQHNCJBOTEHDT63OJCB45BLQWUBPHFVDZZVWJ7TOM
const sourceAccount string = "SDOIUOWJRHVEJRGGTGWYS37HXKU2GKP4KYA4J4M46Z6IRACHJLYK7WMH"

func TxWithTimeBound(minTime, maxTime uint64, a account.Account) (*build.TransactionBuilder, error) {
	tx, err := build.Transaction(
		build.SourceAccount{AddressOrSeed: a.Address()},
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.TestNetwork,
		build.Timebounds{MinTime: minTime, MaxTime: maxTime},
		build.SetOptions()) // do nothing
	if err != nil {
		return nil, err
	}
	return tx, nil
}
