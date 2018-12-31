package blockchain

import (
	"encoding/json"
	"errors"

	"github.com/xoreo/go-basics/blockchain/types/block"
)

// ErrInvalidBlockchain - Error for an attempt to create a new blockchain with invalid parameters
var ErrInvalidBlockchain = errors.New("invalid parameters to construct blockchain")

// ErrNilBlock - Error for a nil block
var ErrNilBlock = errors.New("block is nil")

// Blockchain - A blockchain
type Blockchain struct {
	Blocks []*block.Block `json:"blocks"`
}

// NewBlockchain - Create a new blockchain from a genesis block
func NewBlockchain(genesis *block.Block) (*Blockchain, error) {
	if genesis == nil {
		return nil, ErrInvalidBlockchain
	}

	return &Blockchain{
		Blocks: []*block.Block{genesis},
	}, nil

}

// String - Encode a chain as a string
func (chain *Blockchain) String() string {
	json, _ := json.MarshalIndent(chain, "", "  ")
	return string(json)
}

// AddBlock - Add a block to a blockchain
func (chain *Blockchain) AddBlock(block *block.Block) error {
	if block == nil {
		return ErrNilBlock
	}

	chain.Blocks = append(chain.Blocks, block)
	return nil
}
