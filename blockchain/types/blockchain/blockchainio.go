package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/xoreo/go-basics/blockchain/common"
)

// WriteBlockchainToMemory - Write a blockchain to memory
func (chain *Blockchain) WriteBlockchainToMemory() error {
	json, err := json.MarshalIndent(chain, "", "  ")
	if err != nil {
		return err
	}

	err = common.CreateDirIfDoesNotExist("../../data/chains")
	if err != nil {
		return err
	}

	genesis := chain.Blocks[0]
	hexHash := fmt.Sprintf("%x", genesis.Hash[:8])
	err = ioutil.WriteFile(filepath.FromSlash(fmt.Sprintf("../../data/chains/chain_%s.json", hexHash)), json, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadBlockchainFromMemory - Read a blockchain from memory
func ReadBlockchainFromMemory(hash string) (*Blockchain, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("../../data/chains/chain_%s.json", hash))
	if err != nil {
		return nil, err
	}
	buffer := &Blockchain{}
	err = json.Unmarshal(data, buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
