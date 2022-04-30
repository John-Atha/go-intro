package main

import "fmt"

func add(x int, y int) int {
	return x+y
}

func main() {
	var z int = add(2,3)
	fmt.Println("Result:", z)
}
