package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections on port 1234
	ln, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	conn, err := ln.Accept()
	for {
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
		fmt.Printf("%s\n", buf)
	}

}
