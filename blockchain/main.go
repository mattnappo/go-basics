package main

import (
	"flag"
	"fmt"

	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

var (
	populateFlag = flag.Bool("populate", false, "Populate a blockchain")
)

func main() {
	flag.Parse()
	if *populateFlag {
		fmt.Println("working test")
		chain, err := blockchain.PopulateChain()
		if err != nil {
			panic(err)
		}
		fmt.Println(chain.WriteBlockchainToMemory())
	}
}
