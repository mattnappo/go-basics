package types

import "errors"

// ErrInvalidBlockchain - Error for an attempt to create a new blockchain with invalid parameters
var ErrInvalidBlockchain = errors.New("invalid parameters to construct blockchain")

// ErrNilBlock - Error for a nil block
var ErrNilBlock = errors.New("block is nil")

// Blockchain - A blockchain
type Blockchain struct {
	Blocks []*Block `json:"Blocks"`
}

// NewBlockchain - Create a new blockchain from a genesis block
func NewBlockchain(genesis *Block) (*Blockchain, error) {
	if genesis == nil {
		return nil, ErrInvalidBlockchain
	}

	return &Blockchain{
		Blocks: []*Block{genesis},
	}, nil
	
}

// AddBlock - Add a block to a blockchain
func (blockchain *Blockchain) AddBlock(block *Block) error {
	if block == nil {
		return ErrNilBlock
	}
	
	blockchain.Blocks = append(blockchain.Blocks, block)
	return nil
}