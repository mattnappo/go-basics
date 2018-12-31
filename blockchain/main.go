package main

import (
	"flag"
	"fmt"
)

var (
	populateFlag = flag.Bool("populateChain", false, "Populate a blockchain")
)

func main() {
	if *populateFlag {
		chain, err := types.blockchain.PopulateChain()
		if err != nil {
			panic(err)
		}
		fmt.Println(chain.String())
	}
}
