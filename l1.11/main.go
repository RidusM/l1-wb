package main

import "fmt"

func intersection(a, b []int) []int {
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}

	var res []int

	for _, v := range b {
		if _, ok := m[v]; ok {
			res = append(res, v)
			delete(m, v)
		}
	}

	return res
}

func main() {
	sliceA := []int{1, 2, 3, 5, 8}
	sliceB := []int{2, 3, 4, 8, 10}

	intersect := intersection(sliceA, sliceB)

	fmt.Printf("slice A: %v\n", sliceA)
	fmt.Printf("slice B: %v\n", sliceB)
	fmt.Printf("intersection: %v\n", intersect)
}