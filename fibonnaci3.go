package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		c <- x
	}
}

func main() {
	c := make(chan int)
	go fibonacci(10, c)
	for i := range c {
		fmt.Println(i)
	}
}
