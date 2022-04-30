package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2}
	arr[3] = 18
	fmt.Println(arr, arr[4], len(arr))

	var arr2 [4][4]int
	arr2[0][0] = 10
	arr2[0][3] = 10
	fmt.Println(arr2, len(arr2), arr2[0])

}