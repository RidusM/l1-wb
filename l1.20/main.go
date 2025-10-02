package main

import (
	"fmt"
	"strings"
)

func reverseWordsSimple(s string) string {
    words := strings.Fields(s)
    for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}

func main() {
    testArrays := []string{
        "Reverse Engineering",
		"War.... War never changes",
		"I wanna be hero",
    }
    
    for _, val := range testArrays {
        fmt.Printf("Input: %v\n", val)
        fmt.Printf("Output: %v\n\n", reverseWordsSimple(val))
    }
}