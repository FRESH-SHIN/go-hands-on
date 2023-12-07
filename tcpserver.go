package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Read incoming data
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Close()
		}
		fmt.Printf("Received: %s", buf)
	}

}
