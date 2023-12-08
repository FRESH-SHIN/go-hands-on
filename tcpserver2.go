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
	go sendMessage(conn)
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
func sendMessage(conn net.Conn) {
	var name string
	fmt.Print("your name:")
	fmt.Scan(&name)
	// Send some data to the server
	for true {
		var chat string
		fmt.Scan(&chat)
		bs := fmt.Sprintf("%s:%s", name, chat)
		_, err := conn.Write([]byte(bs))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
