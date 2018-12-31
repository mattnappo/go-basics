package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xoreo/go-basics/blockchain/common"
	"github.com/xoreo/go-basics/blockchain/types/block"
)

// ErrInvalidBlockchain - Error for an attempt to create a new blockchain with invalid parameters
var ErrInvalidBlockchain = errors.New("invalid parameters to construct blockchain")

// ErrValidateIndex - Validation error for when the two blocks' indexes don't match up
var ErrValidateIndex = errors.New("validation error - invalid block indexes")

// ErrValidatePrevHash - Validation error for when the two blocks' hashes don't match up
var ErrValidatePrevHash = errors.New("validation error - hash in block does not match previous block")

// ErrValidateCalcHash - Validation error for when the two blocks' hashes don't match up
var ErrValidateCalcHash = errors.New("validation error - hash in block is not equal to hash of block")

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

// Bytes - Encode a chain as a []byte
func (chain *Blockchain) Bytes() []byte {
	json, _ := json.MarshalIndent(chain, "", "  ")
	return json
}

// AddBlock - Add a block to a blockchain
func (chain *Blockchain) AddBlock(block *block.Block) error {
	if block == nil {
		return ErrNilBlock
	}

	chain.Blocks = append(chain.Blocks, block)
	return nil
}

// Validate - Validate that a block (thisBlock) is valid
func Validate(thisBlock *block.Block, prevBlock *block.Block) error {
	if thisBlock.Index != prevBlock.Index+1 {
		return ErrValidateIndex
	}

	if string(thisBlock.PrevHash) != string(prevBlock.Hash) {
		return ErrValidatePrevHash
	}

	storeHash := thisBlock.Hash
	hashInBlock := fmt.Sprintf("%x", thisBlock.Hash)
	thisBlock.Hash = nil

	fmt.Println(thisBlock.String())

	hashOfBlock := fmt.Sprintf("%x", common.Sha3(thisBlock.Bytes()))

	if hashInBlock != hashOfBlock {
		return ErrValidateCalcHash
	}
	thisBlock.Hash = storeHash

	return nil
}

// ValidateChain - Validate that a chain is valid
func (chain *Blockchain) ValidateChain() error {
	for i := 1; i < len(chain.Blocks); i++ {
		err := Validate(chain.Blocks[i], chain.Blocks[i-1])
		if err != nil {
			return err
		}
	}
	return nil
}
