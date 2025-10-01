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
	fmt.Println(hasUniqueChars("abcd"))
	fmt.Println(hasUniqueChars("abCdefAaf"))
	fmt.Println(hasUniqueChars("aabcd"))
	fmt.Println(hasUniqueChars("abAB"))
	fmt.Println(hasUniqueChars("asdfghjkl"))
	fmt.Println(hasUniqueChars("a"))
}