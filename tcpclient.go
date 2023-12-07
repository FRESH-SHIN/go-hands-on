package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send some data to the server

	_, err = conn.Write([]byte("Hello, server!\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the connection
	conn.Close()
}
