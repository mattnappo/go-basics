package main

import (
	"encoding/json"
	"fmt"

	"github.com/xoreo/go-basics/blockchain/types/blockchain"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
	"github.com/xoreo/go-basics/blockchain/common/util"
)

// Testingchans - test some chans
func Testingchans() {

	// make(chan, int)
	// var myChan chan *blockchain.Blockchain
	myChan := make(chan *blockchain.Blockchain, 1)
	chain := util.GetRandomChain(2)
	myChan <- chain

	tempChain := <-myChan
	fmt.Println(tempChain.String())
}

// Getraw - get raw
func Getraw() (string, error) {
	// Make a slice of transaction pointers
	txns := []*transaction.Transaction{}
	for i := 1; i < 5; i++ {
		txn, err := transaction.NewTransaction(
			"alice",
			"bob",
			20,
		)
		if err != nil {
			return "", err
		}
		txns = append(txns, txn)
	}

	marshalled, err := json.MarshalIndent(txns, "", "  ")
	if err != nil {
		return "", err
	}

	return string(marshalled), nil
}
