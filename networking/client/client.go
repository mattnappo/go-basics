package client

import "net"

// SendMessage - send a message
func SendMessage(s string, conn net.Conn) error {

	// Cast to a byte array
	_, err := conn.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil

}
