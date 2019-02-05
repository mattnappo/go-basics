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
	// fmt.Println(cChain.String())
	// Wait for the client to send their transactions

	// readBuffer := []byte{} // This doesn't really work
	readBuffer := make([]byte, 1000000)                     // 1 KiB
	txnsBuffer := make([]*transaction.Transaction, 1000000) // Transaction buffer for deserialization

	buffer := &transaction.Transaction{}

	for {
		size, err := conn.Read(readBuffer) // Read into buffer
		// fmt.Println(string(buffer))
		if err != nil || size <= 0 {
			return err
		}
		fmt.Printf("read buffer: %sdone\n", string(readBuffer))

		// Unmarshal the data
		// noNewlineReadBuffer := string(readBuffer)
		// noNewlineReadBuffer = strings.TrimSuffix(noNewlineReadBuffer, "\n")

		// err = json.Unmarshal(readBuffer, &txnsBuffer) // Unmarshal the bytes into a slice of transaction pointers

		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		err = json.Unmarshal(readBuffer, buffer)
		if err != nil {
			return err
		}
		fmt.Println("txn buffer: ")
		fmt.Println(txnsBuffer)
		break
	}

	newBlock, err := block.NewBlock(
		cChain.Blocks[len(cChain.Blocks)-1].Index,
		txnsBuffer,
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
