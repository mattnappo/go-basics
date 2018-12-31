package types

import "testing"

func TestNewBlock(t *testing.T) {
	tx1, err := NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	if err != nil {
		t.Fatal(err)
	}

	transactions := []*Transaction{tx1, tx1}

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
