package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

// GetClientTxns - Get the data for a new block from the client's input
func GetClientTxns() ([]*transaction.Transaction, error) {
	reader := bufio.NewReader(os.Stdin)
	var txns []*transaction.Transaction

	fmt.Print("Number of txns to input: ")
	rawNumTxns, _ := reader.ReadString('\n')
	numTxns, err := strconv.Atoi(strings.TrimSuffix(rawNumTxns, "\n"))
	if err != nil {
		return nil, err
	}
	for i := 0; i < numTxns; i++ {
		fmt.Println("==============================")
		fmt.Printf("TRANSACTION %d\n", i)
		fmt.Println("==============================")
		fmt.Print("address to: ")
		to, _ := reader.ReadString('\n')
		to = strings.TrimSuffix(to, "\n")

		fmt.Print("address from: ")
		from, _ := reader.ReadString('\n')
		from = strings.TrimSuffix(from, "\n")

		fmt.Print("amount: ")
		rawAmount, _ := reader.ReadString('\n')
		rawAmount = strings.TrimSuffix(rawAmount, "\n")

		amount, err := strconv.ParseFloat(rawAmount, 64)
		if err != nil {
			return nil, err
		}

		// If any of the inputs are empty, return error
		if to == "" || from == "" || amount == 0.0 {
			return nil, err
		}

		txn, err := transaction.NewTransaction(from, to, amount)
		if err != nil {
			return nil, err
		}
		txns = append(txns, txn)
	}
	return txns, nil
}
