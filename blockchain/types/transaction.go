package types

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
)

// ErrInvalidTransaction - Error for an attempt to create a new transaction with invalid parameters
var ErrInvalidTransaction = errors.New("invalid parameters to construct transaction")

// Transaction - A transaction in a block's list of transactions
type Transaction struct {
	Sender    string  `json:"Sender"`
	Recipient string  `json:"Recipient"`
	Amount    float64 `json:"amount"`
}

// NewTransaction - Create a new transaction
func NewTransaction(sender string, recipient string, amount float64) (*Transaction, error) {
	if sender == "" || recipient == "" || amount == 0 {
		return nil, ErrInvalidTransaction
	}

	return &Transaction{
		sender,
		recipient,
		amount,
	}, nil
}

// String - Encode a transaction into a string
func (transaction *Transaction) String() string {
	json, _ := json.MarshalIndent(*transaction, "", "  ")
	return string(json)
}

// Bytes - Encode a transaction into a []byte
func (transaction *Transaction) Bytes() []byte {
	json, _ := json.MarshalIndent(*transaction, "", "  ")
	return json
}

// TransactionFromBytes - Create a transaction from a []byt
func TransactionFromBytes(b []byte) (*Transaction, error) {
	buffer := &Transaction{}
	err := json.Unmarshal(b, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// WriteTransactionToMemory - Write a transaction to memory
func (transaction *Transaction) WriteTransactionToMemory() error {
	json, err := json.MarshalIndent(*transaction, "", "  ")
	if err != nil {
		return err
	}
	num := big.NewInt(3000)
	seed, err := rand.Int(rand.Reader, num)

	err = ioutil.WriteFile(fmt.Sprintf("transaction_%s.json", seed.String()), json, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadTransactionFromMemory - Read a transaction from memory
func ReadTransactionFromMemory(name string) (*Transaction, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("transaction%s.json", name))
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
