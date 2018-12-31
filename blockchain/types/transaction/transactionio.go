package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/xoreo/go-basics/blockchain/common"
)

// WriteTransactionToMemory - Write a transaction to memory
func (transaction *Transaction) WriteTransactionToMemory() error {
	json, err := json.MarshalIndent(*transaction, "", "  ")
	if err != nil {
		return err
	}

	err = common.CreateDirIfDoesNotExist("data/transactions")
	if err != nil {
		return err
	}

	hexHash := fmt.Sprintf("%x", transaction.Hash[:8])
	err = ioutil.WriteFile(filepath.FromSlash(fmt.Sprintf("data/transactions/transaction_%s.json", hexHash)), json, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadTransactionFromMemory - Read a transaction from memory
func ReadTransactionFromMemory(hash string) (*Transaction, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("data/transactions/transaction_%s.json", hash))
	if err != nil {
		return nil, err
	}
	buffer := &Transaction{}
	err = json.Unmarshal(data, buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
