package networking

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// serverChannel - The main channel for the connections
var serverChannel chan *blockchain.Blockchain

// BroadcastChain - Broadcast the chain to the entire network
func BroadcastChain(conn net.Conn, chain *blockchain.Blockchain) error {
	for {
		time.Sleep(30 * time.Second)
		marshal := chain.String()
		io.WriteString(conn, marshal)
	}
}

// GetGenesis - Get a genesis block
func GetGenesis() *block.Block {
	txns := make([]*transaction.Transaction, 16)
	prev := make([]byte, 16)
	genesis, _ := block.NewBlock(0, txns, prev)
	return genesis
}

// InitServer - Initialize a blockchain server
func InitServer(port string, chain *blockchain.Blockchain) error {
	err := chain.ValidateChain()
	if err != nil {
		return err
	}
	serverChannel = make(chan *blockchain.Blockchain)

	// serverChannel <- chain

	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	defer server.Close()
	for {
		fmt.Println("made chan")
		conn, err := server.Accept()
		if err != nil {
			return err
		}
		go HandleConnection(conn, serverChannel)
		fmt.Println("after")
		// Clients are being accepted and this print statement is running
	}
}

// InitClient - Initialize a client connection
func InitClient(addr string, port string) error {

	conn, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		return err
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	return nil
}
