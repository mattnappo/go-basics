package networking

import (
	"net"
	"os"

	"github.com/xoreo/go-basics/blockchain/types/transaction"

	"github.com/joho/godotenv"
	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"
)

// ServerChannel - The main channel for the connections
var ServerChannel chan *blockchain.Blockchain

// getGenesis - Get a genesis block
func getGenesis() *block.Block {
	txns := make([]*transaction.Transaction, 16)
	prev := make([]byte, 16)
	genesis, _ := block.NewBlock(0, txns, prev)
	return genesis
}

// InitServer - Initialize a blockchain server
func InitServer(chain *blockchain.Blockchain) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = chain.ValidateChain()
	if err != nil {
		return err
	}
	ServerChannel <- chain

	server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		return err
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			return err
		}
		go HandleConnection(conn)
	}
}
