package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/xoreo/go-basics/blockchain/networking"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

var (
	populateFlag = flag.Int("populate", 0, "Populate a blockchain with x blocks")
	loadFlag     = flag.String("load", "", "Load and validate a blockchain from memory")
	serverFlag   = flag.String("server", "6060", "Start a server to host client connections")
	clientFlag   = flag.String("client", "localhost:6060", "Start a client to add blocks to the chain")
)

func main() {
	flag.Parse()
	if *populateFlag > 0 {
		chain, err := blockchain.PopulateChain(*populateFlag)
		if err != nil {
			panic(err)
		}
		chain.WriteBlockchainToMemory()
	}
	if *loadFlag != "" {
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
	if *serverFlag != "" {
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
		params := strings.Split(*clientFlag, ":")
		ip, port := params[0], params[1]
		fmt.Println("ip: " + ip + "\nport: " + port)
		err := networking.InitClient(ip, port)
		if err != nil {
			panic(err)
		}
	}

}
