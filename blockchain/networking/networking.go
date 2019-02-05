package networking

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/xoreo/go-basics/blockchain/common/util"
	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// serverChannel - The main channel for the connections
// var serverChannel chan *blockchain.Blockchain

// BroadcastChain - Broadcast the chain to the entire network
func BroadcastChain(conn net.Conn, chain *blockchain.Blockchain) error {
	for {
		time.Sleep(30 * time.Second)
		conn.Write(chain.Bytes())
		// OR (doesn't matter)
		// marshal := chain.String()
		// io.WriteString(conn, marshal)
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
	serverChannel := make(chan *blockchain.Blockchain, 1)
	serverChannel <- chain

	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	fmt.Println("server initialized")
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			return err
		}
		go HandleConnection(conn, serverChannel)
		// Clients are being accepted and this print statement is running
	}
}

// InitClient - Initialize a client connection
func InitClient(addr string, port string) error {
	// Connect to the server
	conn, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		return err
	}
	fmt.Println("connected to server node ")

	// Get transactions
	txns, err := util.GetClientTxns()
	if err != nil {
		return err
	}
	for _, txn := range txns {
		fmt.Print(txn.String())
	}

	// Write transactions to server
	bytes, _ := json.MarshalIndent(txns, "", "  ")
	// Write to file for testing purposes
	// err = ioutil.WriteFile("save.json", bytes, 0644)
	conn.Write(bytes)
	return nil
}
