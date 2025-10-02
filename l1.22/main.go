package main

import (
	"fmt"
	"math/big"
)

func method1WithInt() {
	var a int = 2_000_000
	var b int = 5_000_000

	fmt.Printf("a = %d, b = %d\n\n", a, b)

	sum := a + b
	fmt.Printf("a + b: %d\n", sum)

	sub := b - a
	fmt.Printf("b - a: %d\n", sub)

	mul := a * b
	fmt.Printf("a * b: %d\n", mul)

	div := b / a
	fmt.Printf("b / a: %d\n", div)
	
	fmt.Println()
}

func method2WithBigInt() {
	aStr := "20000000000000000000"
	bStr := "50000000000000000000" 

	a, okA := new(big.Int).SetString(aStr, 10)
	b, okB := new(big.Int).SetString(bStr, 10)

	if !okA || !okB {
		fmt.Println("error while convert string to big.Int")
		return
	}
	
	fmt.Printf("a = %s\nb = %s\n\n", a.String(), b.String())

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("a + b: %s\n", sum.String())

	sub := new(big.Int)
	sub.Sub(b, a)
	fmt.Printf("b - a: %s\n", sub.String())

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Printf("a * b: %s\n", mul.String())

	div := new(big.Int)
	div.Div(b, a)
	fmt.Printf("b / a: %s\n", div.String())
	
	fmt.Println()
}

func main() {
	fmt.Println("Method 1: with Integer")
	method1WithInt()
	fmt.Println("Method 2: with Big Integer")
	method2WithBigInt()
}