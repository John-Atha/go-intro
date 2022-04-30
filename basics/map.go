package main

import "fmt"

func has(dict map[string]int, key string) bool {
	_, prs := dict[key]
	fmt.Println(key, " in ", dict, ": ", prs)
	return prs
}

func keys(dict map[string]int) []string {
	keys_ := make([]string, 0)
	for key := range dict {
		keys_ = append(keys_, key)
	}
	fmt.Println("Keys of ", dict, ": ", keys_)
	return keys_
}

func values(dict map[string]int) []int {
	vals := make([]int, 0)
	for _, val := range dict {
		vals = append(vals, val)
	}
	fmt.Println("Values of ", dict, ": ", vals)
	return vals
}

func main() {
	m := make(map[string]int)

	m["a"] = 5
	m["b"] = 1
	m["c"] = 3
	m["d"] = 8

	has(m, "a")
	has(m, "2")
	has(m, "f")
	has(m, "d")

	keys(m)
	values(m)
}