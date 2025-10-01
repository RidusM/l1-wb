package main

import "fmt"

func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

func main() {
    sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    testCases := []int{
        1,
        19,
        7,
        10,
        0,
        20,
    }
    
    for _, target := range testCases {
        index := binarySearch(sortedArray, target)
        if index != -1 {
            fmt.Printf("Element %d found by index %d\n", target, index)
        } else {
            fmt.Printf("Element %d not found: %d \n", target, index)
        }
    }
}