package main

import "fmt"

func sum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	go sum(1, 3)
	go sum(2, 3)
}
