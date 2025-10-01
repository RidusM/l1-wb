package main

import "fmt"

func main() {
	a1, b1 := 5, 10

	// ARITHMETIC
	fmt.Printf("a = %d, b = %d\n", a1, b1)
	a1 = a1 + b1
	b1 = a1 - b1
	a1 = a1 - b1
	fmt.Printf("a = %d, b = %d\n\n", a1, b1)

	// XOR
	a2, b2 := 5, 10
	fmt.Printf("a = %d, b = %d\n", a2, b2)
	a2 = a2 ^ b2
	b2 = a2 ^ b2
	a2 = a2 ^ b2
	fmt.Printf("a = %d, b = %d\n\n", a2, b2)

	// GO 
	a3, b3 := 5, 10
	fmt.Printf("a = %d, b = %d\n", a3, b3)
	a3, b3 = b3, a3
	fmt.Printf("a = %d, b = %d\n", a3, b3)
}