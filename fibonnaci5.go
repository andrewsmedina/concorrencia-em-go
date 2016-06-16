package main

import (
	"fmt"
)

func fibonacci(c chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
	}
}

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-c)
		}
	}()
	fibonacci(c)
}
