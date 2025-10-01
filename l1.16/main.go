package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    pivot := arr[len(arr)/2]

    left := make([]int, 0)
    middle := make([]int, 0)
    right := make([]int, 0)
    
    for _, item := range arr {
        switch {
        case item < pivot:
            left = append(left, item)
        case item == pivot:
            middle = append(middle, item)
        case item > pivot:
            right = append(right, item)
        }
    }
    
    return append(append(quickSort(left), middle...), quickSort(right)...)
}

func main() {
    testArrays := [][]int{
        {64, 34, 25, 12, 22, 11, 90},
        {5, 2, 4, 6, 1, 3},
        {1},
        {},
        {3, 3, 3, 3},
        {9, 8, 7, 6, 5, 4, 3, 2, 1},
    }
    
    for _, arr := range testArrays {
        fmt.Printf("Input: %v\n", arr)
        fmt.Printf("Sort: %v\n\n", quickSort(arr))
    }
}