package main

import (
	"fmt"
)

func clearBit(num int64, position uint) int64 {
	mask := int64(1) << position
	return num &^ mask
}

func main() {
	testCases := [][]int64{
        {5, 0},
		{7, 1},
		{100, 5},
    }
    
    for _, val := range testCases {
        fmt.Printf("Input: %v, %v\n", val[0], val[1])
        fmt.Printf("Output: %v\n\n", clearBit(val[0], uint(val[1])))
    }
}