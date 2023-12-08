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
	channels := []chan string{}
	receive_channel := make(chan string)
	go broadcast(receive_channel, &channels)
	go serverMessage(receive_channel)
	// Accept incoming connections and handle them
	for true {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		strchan := make(chan string)
		channels = append(channels, strchan)
		go receiveMessage(conn, receive_channel)
		go sendMessage(conn, strchan)
	}

}

func serverMessage(receive_channel chan string) {
	var name string
	fmt.Print("your name:")
	fmt.Scan(&name)
	// Send some data to the server
	for true {
		var chat string
		fmt.Scan(&chat)
		bs := fmt.Sprintf("%s:%s", name, chat)
		receive_channel <- bs
	}
}
func broadcast(receive_channel chan string, channels *[]chan string) {
	for true {
		str := <-receive_channel
		fmt.Printf("%s", *channels)
		for _, c := range *channels {
			c <- str
		}
	}
}
func sendMessage(conn net.Conn, c chan string) {
	// Send some data to the server
	for true {
		str := <-c
		fmt.Println("received from channel")
		_, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
func receiveMessage(conn net.Conn, receive_channel chan string) {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			conn.Close()
		}
		fmt.Printf("%s\n", buf)
		receive_channel <- fmt.Sprintf("%s", buf)
	}
}
