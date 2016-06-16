package main

import (
	"fmt"
	"time"
)

func sum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	go sum(1, 3)
	go sum(2, 3)
	time.Sleep(1 * time.Second)
}
