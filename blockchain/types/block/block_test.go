package block

import (
	"testing"

	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

func TestNewBlock(t *testing.T) {
	tx1, err := transaction.NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	if err != nil {
		t.Fatal(err)
	}

	transactions := []*transaction.Transaction{tx1, tx1}

	block, err := NewBlock(
		0,
		transactions,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(block)
}
