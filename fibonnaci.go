package main

import (
	"fmt"
)

func fibonacci(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

func main() {
	fibonacci(10)
}
