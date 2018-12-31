package blockchain

import (
	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// createTransactions - Create some nonsense txs
func createTransactions() []*transaction.Transaction {
	tx1, _ := transaction.NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	tx2, _ := transaction.NewTransaction(
		"matt",
		"dowland",
		0.78,
	)

	transactions := []*transaction.Transaction{
		tx1,
		tx2,
	}
	return transactions
}

func newBlock() (*block.Block, error) {
	// Get some (nonsense) transactions
	txs := createTransactions()
	// Create the genesis block
	newBlock, err := block.NewBlock(
		0,
		txs,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return newBlock, nil
}

func prevBlock(lastBlock block.Block) (*block.Block, error) {
	prevBlock, err := block.NewBlock(
		lastBlock.Index+1,
		createTransactions(),
		lastBlock.Hash,
	)
	if err != nil {
		return nil, err
	}
	return prevBlock, nil
}

// PopulateChain - Populate a chain with nonsense blocks/txs
func PopulateChain() (*Blockchain, error) {
	// Get some (nonsense) transactions
	txs := createTransactions()

	// Create the genesis block
	genesis, err := block.NewBlock(
		0,
		txs,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Create a chain
	chain, err := NewBlockchain(genesis)
	if err != nil {
		return nil, err
	}

	currentBlock := genesis
	for i := 0; i < 10; i++ {
		if err != nil {
			return nil, err
		}
		newBlock, err := prevBlock(*currentBlock)
		if err != nil {
			return nil, err
		}
		chain.AddBlock(newBlock)
		currentBlock = newBlock
	}

	return chain, nil
}
