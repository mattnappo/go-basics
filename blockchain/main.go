package main

import (
	"flag"

	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

var (
	populateFlag = flag.Bool("populate", false, "Populate a blockchain")
	// find ./ -type d -name data -exec rm -rf {} \\;
)

func main() {
	flag.Parse()
	if *populateFlag {
		chain, err := blockchain.PopulateChain()
		if err != nil {
			panic(err)
		}
		chain.WriteBlockchainToMemory()
	}

}
