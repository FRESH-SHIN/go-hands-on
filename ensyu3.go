package main

import (
	"fmt"
	"time"
)

func fizz(c chan string) {
	for true {
		time.Sleep(3000 * time.Millisecond)
		c <- "fizz"
	}
}
func buzz(c chan string) {
	for true {
		time.Sleep(5000 * time.Millisecond)
		c <- "buzz"
	}
}

func main() {
	c := make(chan string, 10)
	go fizz(c)
	go buzz(c)
	time.Sleep(1000 * time.Millisecond)
	for i := 1; ; i++ {
		fmt.Printf("%d : ", i)
		for len(c) > 0 {
			fmt.Print(<-c)
		}
		fmt.Print("\n")
		time.Sleep(1000 * time.Millisecond)
	}
}
