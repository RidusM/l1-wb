package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})

	for _, item := range data {
		m[item] = struct{}{}
	}

	res := make([]string, 0, len(m))
	for key := range m {
		res = append(res, key)
	}
	
	fmt.Printf("iput slice: %v\n", data)
	fmt.Printf("resulting slice: %v\n", res)
}