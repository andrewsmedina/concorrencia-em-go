select {
case c <- x:
	x, y = y, x+y
case <-quit:
	fmt.Println("quit")
	return
}
