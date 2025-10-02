package main

import "fmt"

func reverseIterative(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	testArrays := []string{
        "Reverse Engineering",
		"War.... War never changes",
		"I wanna be hero",
    }
    
    for _, val := range testArrays {
        fmt.Printf("Input: %v\n", val)
        fmt.Printf("Output: %v\n\n", reverseIterative(val))
    }
}