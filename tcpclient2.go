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
	go receiveMessage(conn)
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
	// Close the connection
	conn.Close()
}

func receiveMessage(conn net.Conn) {
	for true {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Close()
		}
		fmt.Printf("%s\n", buf)
	}

}
