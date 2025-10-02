package main

import "fmt"

func main() {
	a1, b1 := 5, 10

	fmt.Println("Method 1: Arithmetic")
	fmt.Printf("a = %d, b = %d\n", a1, b1)
	a1 = a1 + b1
	b1 = a1 - b1
	a1 = a1 - b1
	fmt.Printf("a = %d, b = %d\n\n", a1, b1)

	fmt.Println("Method 2: XOR")
	a2, b2 := 5, 10
	fmt.Printf("a = %d, b = %d\n", a2, b2)
	a2 = a2 ^ b2
	b2 = a2 ^ b2
	a2 = a2 ^ b2
	fmt.Printf("a = %d, b = %d\n\n", a2, b2)

	fmt.Println("Method 3: GO")
	a3, b3 := 5, 10
	fmt.Printf("a = %d, b = %d\n", a3, b3)
	a3, b3 = b3, a3
	fmt.Printf("a = %d, b = %d\n", a3, b3)
}