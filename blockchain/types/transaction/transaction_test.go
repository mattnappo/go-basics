package transaction

import "testing"

func TestNewTransaction(t *testing.T) {
	tx, err := NewTransaction(
		"dowland",
		"matt",
		10000000,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)
}
