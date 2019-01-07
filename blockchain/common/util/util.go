package util

import (
	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// GetRandomChain - Get a random blockchain with nonsense blocks & txns
func GetRandomChain(size int) *blockchain.Blockchain {
	tx1, _ := transaction.NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	transactions := []*transaction.Transaction{tx1, tx1}
	genesis, _ := block.NewBlock(
		0,
		transactions,
		nil,
	)
	blockchain, _ := blockchain.NewBlockchain(genesis)
	newBlock, _ := block.NewBlock(
		genesis.Index+1,
		transactions,
		genesis.Hash,
	)
	for i := 0; i < size-1; i++ {
		blockchain.AddBlock(newBlock)
	}

	_ = blockchain.ValidateChain()
	return blockchain
}
