package main

import (
	"fmt"
	"time"
)

func fizz() {
	for true {
		time.Sleep(3000 * time.Millisecond)
		fmt.Print("fizz")
	}
}
func buzz() {
	for true {
		time.Sleep(5000 * time.Millisecond)
		fmt.Print("buzz")
	}
}

func main() {
	go fizz()
	go buzz()
	for i := 1; ; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%d : \n", i)
	}
}
