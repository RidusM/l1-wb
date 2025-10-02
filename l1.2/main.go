package main

import (
	"fmt"
	"sync"
)

func main() {
	in := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for _, val := range in {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("%d\n", num*num)
		}(val)
	}
	wg.Wait()
}