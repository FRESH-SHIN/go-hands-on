package main

import (
	"fmt"
	"time"
)

func say(c chan int) {
	for i := 0; ; i++ {
		time.Sleep(1000 * time.Millisecond)
		c <- i
	}
}

func main() {
	c := make(chan int, 10)
	go say(c)
	var receivedInt int
	for true {
		receivedInt = <-c
		fmt.Printf("受け取ったデータ：%d\n", receivedInt)
	}
}
