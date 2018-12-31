package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/xoreo/go-basics/blockchain/common"
)

// WriteBlockToMemory - Write a block to memory
func (block *Block) WriteBlockToMemory() error {
	json, err := json.MarshalIndent(*block, "", "  ")
	if err != nil {
		return err
	}

	err = common.CreateDirIfDoesNotExist("../../data/blocks")
	if err != nil {
		return err
	}

	hexHash := fmt.Sprintf("%x", block.Hash[:8])
	err = ioutil.WriteFile(filepath.FromSlash(fmt.Sprintf("../../data/blocks/block_%s.json", hexHash)), json, 0644)
	if err != nil {
		return err
	}
	return nil

}

// ReadBlockFromMemory - Read a block from memory
func ReadBlockFromMemory(hash string) (*Block, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("../../data/blocks/block_%s.json", hash))
	if err != nil {
		return nil, err
	}
	buffer := &Block{}
	err = json.Unmarshal(data, buffer)
	// buffer.Timestamp = data.Timestamp
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
