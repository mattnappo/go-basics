package transaction

import (
	"fmt"
	"testing"
)

func TestWriteTransactionToMemory(t *testing.T) {
	tx, err := NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = tx.WriteTransactionToMemory()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)
}

func TestReadTransactionFromMemory(t *testing.T) {
	tx, err := NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = tx.WriteTransactionToMemory()
	if err != nil {
		t.Fatal(err)
	}

	hexHash := fmt.Sprintf("%x", tx.Hash)[:8]
	tx, err = ReadTransactionFromMemory(string(hexHash))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)
}
