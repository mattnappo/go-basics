package networking

import (
	"fmt"
	"net"

	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

// HandleConnection - Handle a connection with a client
func HandleConnection(conn net.Conn, channel chan *blockchain.Blockchain) {
	defer conn.Close()
	// cChain := <-channel
	// txns := nil
	// Wait for the client to send their transactions
	buffer := make([]byte, 1000)
	for {
		size, err := conn.Read(buffer) // Read into buffer
		if err == nil && size > 0 {    // Check for errors
			fmt.Printf("\ntyped: %s\n", string(buffer))

			// data, err := ioutil.ReadFile(fmt.Sprintf("data/blocks/block_%s.json", hash))
			// if err != nil {
			// 	return nil, err
			// }
			// buffer := &Block{}
			// err = json.Unmarshal(data, buffer)
			// if err != nil {
			// 	return nil, err
			// }
			// return buffer, nil

			// txns = buffer
			// break
		}
		buffer = nil
	}

	// newBlock, err := block.NewBlock(
	// 	cChain.Blocks[len(cChain.Blocks)-1].Index,
	// 	txns,
	// 	cChain.Blocks[len(cChain.Blocks)-1].Hash,
	// )
	// if err != nil {
	// 	fmt.Println("error - invalid construction of block")
	// 	conn.Close()
	// }

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
