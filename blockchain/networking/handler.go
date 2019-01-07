package networking

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/xoreo/go-basics/blockchain/types/block"
	"github.com/xoreo/go-basics/blockchain/types/blockchain"

	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// HandleConnection - Handle a connection with a client
func HandleConnection(conn net.Conn, chain chan *blockchain.Blockchain) {
	defer conn.Close()
	var txns []*transaction.Transaction

	io.WriteString(conn, "Number of txns to input: ")
	scanner := bufio.NewScanner(conn)
	go func() {
		for scanner.Scan() {
			txnCount, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("error - not a valid amount of txns to input")
				continue
			}
			for i := 0; i < txnCount; i++ {
				io.WriteString(conn, "to: ")
				to := scanner.Text()

				io.WriteString(conn, "from: ")
				from := scanner.Text()

				io.WriteString(conn, "amount: ")
				amount := scanner.Text()

				f, err := strconv.ParseFloat(amount, 64)
				if err != nil {
					fmt.Println("error - invalid amount field")
					continue
				}
				txn, err := transaction.NewTransaction(from, to, f)
				if err != nil {
					fmt.Println("error - txn creation error")
				}
				txns = append(txns, txn)
			}
			blocks := <-chain.Blocks
			newBlock, err := block.NewBlock()

		}
	}()
}
