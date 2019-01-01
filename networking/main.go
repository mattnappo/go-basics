package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/xoreo/go-basics/networking/client"
	"github.com/xoreo/go-basics/networking/server"
)

var (
	// Flag.Bool returns a pointer
	serverFlag = flag.Bool("server", false, "Starts a server")
	clientFlag = flag.Bool("client", false, "Starts a client")
)

func main() {
	flag.Parse()

	if *serverFlag {
		err := server.StartServer()

		if err != nil { // Check for errors
			panic(err) // Panic
		}
	} else if *clientFlag {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}

		reader := bufio.NewScanner(os.Stdin)
		for {
			fmt.Println("Enter a message:")

			reader.Scan()

			input := reader.Text()
			input = strings.TrimSuffix(input, "\n")

			client.SendMessage(input, conn)
		}
	}
}
