package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

func main() {
	go fibonacci(10)
	time.Sleep(1 * time.Second)
}
