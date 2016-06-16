package main

import (
	"fmt"
)

func fibonacci(n int) {
	x, y := 0
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

func main() {
	go fibonacci(10)
}
