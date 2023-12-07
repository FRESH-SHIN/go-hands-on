package main

import "fmt"

func main() {
	a := "Hello world"
	result := printSomeTexts(a, 100)
	fmt.Printf("%d %s \n", result, "Done!")
}
func printSomeTexts(msg string, b int) int {
	sum := 0
	for i := 1; i <= b; i++ {
		sum = sum + i
	}
	return sum
}
