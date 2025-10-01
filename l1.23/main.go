package main

import "fmt"


func removeSimple(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}


func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("input slice: %v\n", numbers)
	
	numbers = removeSimple(numbers, 2)
	fmt.Printf("after delete [2]: %d\n", numbers)
}