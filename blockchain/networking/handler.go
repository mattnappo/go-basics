package networking

import (
	"encoding/json"
	"net"

	"github.com/xoreo/go-basics/blockchain/types/blockchain"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// HandleConnection - Handle a connection with a client
func HandleConnection(conn net.Conn, channel chan *blockchain.Blockchain) error {
	defer conn.Close()
	// cChain := <-channel

	// Wait for the client to send their transactions

	txns := []*transaction.Transaction{}
	buffer := make([]byte, 1000000) // 1 KiB
	for {
		size, err := conn.Read(buffer) // Read into buffer
		if err != nil || size <= 0 {
			return err
		}

		err = json.Unmarshal(buffer, txns) // Unmarshal the bytes into a slice of transaction pointers
		if err != nil {
			return err
		}
		buffer = nil
		break
	}

	// newBlock, err := block.NewBlock(
	// 	cChain.Blocks[len(cChain.Blocks)-1].Index,
	// 	txns,
	// 	cChain.Blocks[len(cChain.Blocks)-1].Hash,
	// )
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(newBlock.String())
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
