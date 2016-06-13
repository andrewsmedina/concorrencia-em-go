package main

import "fmt"

func hello() {
	fmt.Println("oie")
}

func main() {
	go hello()
}
