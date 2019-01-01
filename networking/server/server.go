package server

import (
	"fmt"
	"net"

	"github.com/xoreo/go-basics/networking/handler"
)

// StartServer - Start a server
func StartServer() error {
	fd, err := net.Listen("tcp", ":8080")

	if err != nil {
		return err
	}

	// for {
	connection, err := fd.Accept()

	if err == nil {
		fmt.Println("Accepted client.")
		handler.Handle(connection)
	}
	// }
	return nil
}
