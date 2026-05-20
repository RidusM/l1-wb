package main

import (
	"fmt"
	"unicode"
)

func hasUniqueChars(s string) bool {
	if len(s) <= 1{
		return true
	}

	seen := make(map[rune]struct{}, len(s))

	for _, char := range s {
		lowerChar := unicode.ToLower(char)
		if _, exists := seen[lowerChar]; exists {
			return false
		}
		seen[lowerChar] = struct{}{}
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
