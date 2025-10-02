package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[rune]bool)

	for _, char := range s {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	testCases := []string{
        "abcd",
		"abCdefAaf",
		"AABCD",
		"adAB",
		"asdfghjkl",
		"a",
		"",
    }
    
    for _, val := range testCases {
        fmt.Printf("Input: %v\n", val)
        fmt.Printf("Output: %v\n\n", hasUniqueChars(val))
    }
}