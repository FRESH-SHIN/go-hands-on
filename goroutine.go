package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for true {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	for true {
		time.Sleep(1000)
	}
}
