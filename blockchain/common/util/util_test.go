package util

import (
	"fmt"
	"testing"
)

func TestGetRandomChain(t *testing.T) {
	chain := GetRandomChain(3)
	fmt.Println(chain.String())
	chain = GetRandomChain(5)
	fmt.Println(chain.String())
}

func TestGetClientTxns(t *testing.T) {
	txns, err := GetClientTxns()
	if err != nil {
		t.Log(err)
	}
	for _, txn := range txns {
		fmt.Print(txn.String())
	}
}
