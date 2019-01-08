package main

import (
	"flag"
	"fmt"

	"github.com/xoreo/go-basics/blockchain/networking"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

var (
	populateFlag = flag.Int("populate", 0, "Populate a blockchain with x blocks")
	loadFlag     = flag.String("load", "", "Load and validate a blockchain from memory")
	generateFlag = flag.Bool("generate", false, "Generate a new block onto the chain")
)

func main() {
	flag.Parse()
	if *populateFlag > 0 {
		fmt.Println("POPULATE FLAG")
		chain, err := blockchain.PopulateChain(*populateFlag)
		if err != nil {
			panic(err)
		}
		chain.WriteBlockchainToMemory()
	}
	if *loadFlag != "" {
		fmt.Println("LOAD FLAG")
		// name := fmt.Sprintf("data/chains/chain_%s.json", *loadFlag)
		chain, err := blockchain.ReadBlockchainFromMemory(*loadFlag)
		if err != nil {
			panic(err)
		}
		err = chain.ValidateChain()
		if err != nil {
			panic(err)
		}
		// fmt.Println(chain)
	}
	if *generateFlag {
		genesis := networking.GetGenesis()
		chain, err := blockchain.NewBlockchain(genesis)
		if err != nil {
			panic(err)
		}
		err = networking.InitServer(chain)
		if err != nil {
			panic(err)
		}

	}

}
