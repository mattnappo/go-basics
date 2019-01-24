package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/xoreo/go-basics/blockchain/common/util"
	"github.com/xoreo/go-basics/blockchain/networking"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

var (
	populateFlag = flag.Int("populate", 0, "Populate a blockchain with x blocks")
	loadFlag     = flag.String("load", "", "Load and validate a blockchain from memory")
	serverFlag   = flag.String("server", "", "Start a server to host client connections")
	clientFlag   = flag.String("client", "", "Start a client to add blocks to the chain")
)

func testingchans() {

	// make(chan, int)
	// var myChan chan *blockchain.Blockchain
	myChan := make(chan *blockchain.Blockchain, 1)
	chain := util.GetRandomChain(2)
	myChan <- chain

	tempChain := <-myChan
	fmt.Println(tempChain.String())
}

func testing() {

}

func main() {
	// networking.TestChans()

	flag.Parse()
	if *populateFlag > 0 {
		// fmt.Println("populate flag")
		chain, err := blockchain.PopulateChain(*populateFlag)
		if err != nil {
			panic(err)
		}
		chain.WriteBlockchainToMemory()
	}

	if *loadFlag != "" {
		// fmt.Println("load flag")
		// name := fmt.Sprintf("data/chains/chain_%s.json", *loadFlag)
		chain, err := blockchain.ReadBlockchainFromMemory(*loadFlag)
		if err != nil {
			panic(err)
		}
		err = chain.ValidateChain()
		if err != nil {
			panic(err)
		}
	}

	if *serverFlag != "" {
		// fmt.Println("server flag")
		genesis := networking.GetGenesis()
		chain, err := blockchain.NewBlockchain(genesis)
		if err != nil {
			panic(err)
		}
		err = networking.InitServer(*serverFlag, chain)

		if err != nil {
			panic(err)
		}
	}

	if *clientFlag != "" {
		// fmt.Println("client flag")
		params := strings.Split(*clientFlag, ":")
		ip, port := params[0], params[1]
		err := networking.InitClient(ip, port)
		if err != nil {
			panic(err)
		}
	}

}
