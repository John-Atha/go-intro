package main

import "fmt"

func remove(arr_init []int, x int) []int {
	// without
	arr := make([]int, len(arr_init))
	copy(arr, arr_init)
	fmt.Print("Removing ", x, " from ", arr)
	length := len(arr)
	i := 0
	for i = 0; i < length; i++ {
		if arr[i] == x {
			break
		}
	}
	var res []int
	if i < length-1 {
		res = append(arr[0:i], arr[(i+1):]...)
	} else if i == length-1 {
		res = arr[:length-1]
	} else {
		res = arr
	}
	fmt.Println(" => ", res)
	return res
}

func findIndex(arr []int, x int) int {
	fmt.Print("Index of ", x, " in ", arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			fmt.Println(" => ", i)
			return i
		}
	}
	fmt.Println(" => ", -1)
	return -1
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	s = remove(s, 1)
	s = remove(s, 2)
	s = remove(s, 4)
	s = remove(s, 1000)
	s = remove(s, 7)

	findIndex(s, 2)
	findIndex(s, 3)
	findIndex(s, 8)
	findIndex(s, 1000)
	findIndex(s, 4)

	fmt.Println("i :  val")
	for index, val := range s {
		fmt.Println(index, ": ", val)
	}

}
