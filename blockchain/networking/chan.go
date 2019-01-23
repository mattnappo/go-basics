package networking

import (
	"fmt"

	"github.com/xoreo/go-basics/blockchain/common/util"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

// TestChans - Test some basic channel concepts
func TestChans() {
	// var myChan chan *blockchain.Blockchain
	myChan := make(chan *blockchain.Blockchain, 1)
	chain := util.GetRandomChain(2)
	myChan <- chain

	tempChain := <-myChan
	fmt.Println(tempChain.String())
}
