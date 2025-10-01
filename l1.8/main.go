package main

import (
	"fmt"
)

func clearBit(num int64, position uint) int64 {
	mask := int64(1) << position
	return num &^ mask
}

func main() {
	fmt.Println(clearBit(5, 0))
	fmt.Println(clearBit(7, 1))
}