package main

import "fmt"

func main() {
	testArrays := [][]string{
		{"cat", "cat", "dog", "cat", "tree"},
		{"zombie", "dwarf", "loxodon", "elf", "loxodon"},
		{"apple", "orange", "apple", "peach", "peach"},
	}

	for _, arr := range testArrays {
		m := make(map[string]struct{})
		
		for _, item := range arr {
			m[item] = struct{}{}
		}

		res := make([]string, 0, len(m))
		for key := range m {
			res = append(res, key)
		}
		
		fmt.Printf("Input slice: %v\n", arr)
		fmt.Printf("Output slice: %v\n\n", res)
	}
}