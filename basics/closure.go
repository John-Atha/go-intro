package main

import "fmt"

func intSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := intSequence()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println("---------------")
	f2 := intSequence()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println("---------------")
	fmt.Println(f())
	fmt.Println("---------------")
	fmt.Println(f2())
}