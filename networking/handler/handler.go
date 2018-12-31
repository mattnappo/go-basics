package handler

import (
	"fmt"
	"net"
)

// Handle - Handle a connection
func Handle(connection net.Conn) {
	buffer := make([]byte, 256)
	for {
		size, err := connection.Read(buffer) // Read into buffer
		if err == nil && size > 0 {          // Check for errors
			fmt.Printf("\ntyped: %s\n", string(buffer))
		}

		// buffer = nil
	}

	// connection.Close()
}
