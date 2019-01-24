package networking

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// HandleConnection - Handle a connection with a client
func HandleConnection(conn net.Conn, channel chan *blockchain.Blockchain) error {
	defer conn.Close()
	cChain := <-channel
	fmt.Println(cChain.String())
	// Wait for the client to send their transactions

	txns := make([]*transaction.Transaction, 10000)
	buffer := make([]byte, 1000000) // 1 KiB
	for {
		size, err := conn.Read(buffer) // Read into buffer
		// fmt.Println(string(buffer))
		if err != nil || size <= 0 {
			return err
		}
		fmt.Println("befor ecreating new lblock")
		fmt.Println(string(buffer))

		err = json.Unmarshal(buffer, &txns) // Unmarshal the bytes into a slice of transaction pointers
		if err != nil {
			fmt.Println("bad")
			fmt.Println(err)
			return err
		}
		fmt.Println("hello world!")
		break
	}

	newBlock, err := block.NewBlock(
		cChain.Blocks[len(cChain.Blocks)-1].Index,
		txns,
		cChain.Blocks[len(cChain.Blocks)-1].Hash,
	)
	if err != nil {
		return err
	}
	fmt.Println("after creating new block: sjowing new block nwo:")
	fmt.Println(newBlock.String())
	return nil

	// err = blockchain.Validate(newBlock, cChain.Blocks[len(cChain.Blocks)-1])
	// if err != nil {
	// 	fmt.Println("error - new block is invalid after checks")
	// 	conn.Close()
	// }
	// err = cChain.AddBlock(newBlock)
	// if err != nil {
	// 	fmt.Println("error - new block could not be added to the chain")
	// 	conn.Close()
	// }
	// channel <- cChain
	// go BroadcastChain(conn, cChain)
	// fmt.Printf("%s", cChain.String())
}
