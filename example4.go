package main

import (
	"fmt"
)

func sum(a, b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int)
	go sum(1, 3, c)
	go sum(2, 3, c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
