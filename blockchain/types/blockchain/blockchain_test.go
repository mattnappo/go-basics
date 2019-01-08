package blockchain

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/xoreo/go-basics/blockchain/common/util"
// 	"github.com/xoreo/go-basics/blockchain/types/block"
// 	"github.com/xoreo/go-basics/blockchain/types/transaction"
// )

// func TestNewBlockchain(t *testing.T) {
// 	tx1, err := transaction.NewTransaction(
// 		"dowland",
// 		"matt",
// 		10000000,
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	transactions := []*transaction.Transaction{tx1, tx1}

// 	genesis, err := block.NewBlock(
// 		0,
// 		transactions,
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(genesis)

// 	blockchain, err := NewBlockchain(genesis)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(blockchain)
// }

// func TestAddBlock(t *testing.T) {
// 	tx1, err := transaction.NewTransaction(
// 		"dowland",
// 		"matt",
// 		10000000,
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	transactions := []*transaction.Transaction{tx1, tx1}

// 	genesis, err := block.NewBlock(
// 		0,
// 		transactions,
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(genesis)

// 	blockchain, err := NewBlockchain(genesis)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(blockchain)
// 	newBlock, err := block.NewBlock(
// 		genesis.Index+1,
// 		transactions,
// 		genesis.Hash,
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(newBlock)
// 	blockchain.AddBlock(newBlock)
// 	t.Log(blockchain)
// }

// func TestValidate(t *testing.T) {
// 	tx1, err := transaction.NewTransaction(
// 		"dowland",
// 		"matt",
// 		10000000,
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	transactions := []*transaction.Transaction{tx1, tx1}

// 	genesis, err := block.NewBlock(
// 		0,
// 		transactions,
// 		nil,
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(genesis)

// 	blockchain, err := NewBlockchain(genesis)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(blockchain)
// 	newBlock, err := block.NewBlock(
// 		genesis.Index+1,
// 		transactions,
// 		genesis.Hash,
// 	)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(newBlock)
// 	blockchain.AddBlock(newBlock)
// 	t.Log(blockchain)

// 	err = Validate(newBlock, genesis)
// 	if err != nil {
// 		t.Fatal(err)
// 	} else {
// 		t.Log("success - valid blocks")
// 	}

// }

// func TestValidateChain(t *testing.T) {
// 	blockchain := util.GetRandomChain(5)

// 	err := blockchain.ValidateChain()
// 	if err != nil {
// 		t.Fatal(err)
// 	} else {
// 		t.Log("success - valid blockchain")
// 	}
// }

// func TestReplaceLongerChain(t *testing.T) {
// 	newChain := util.GetRandomChain(12)
// 	fmt.Printf("newChain size: %d\n", len(newChain.Blocks))
// 	oldChain := util.GetRandomChain(11)
// 	fmt.Printf("oldChain size: %d\n", len(oldChain.Blocks))
// 	replaced := ReplaceLongerChain(newChain, oldChain)
// 	if replaced {
// 		t.Log("success - longer chain replaced")
// 	} else {
// 		t.Fatal("error - chain not replaced")
// 	}
// }
