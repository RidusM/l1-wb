package main

import (
	"fmt"
	"sync"
)

func main() {
	in := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for _, val := range in{
		wg.Go(func(){
			fmt.Printf("%d\n", val*val)
		})
	}
	wg.Wait()
}