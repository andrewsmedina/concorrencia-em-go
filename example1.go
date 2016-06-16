package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(sum(1, 3))
	fmt.Println(sum(2, 3))
}
