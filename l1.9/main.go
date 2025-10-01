package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()

	for val := range in {
		out <- val * 2
	}
	close(out)
}

func printer(wg *sync.WaitGroup, in <-chan int) {
	defer wg.Done()
	for val := range in {
		fmt.Printf("%d\n", val)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	n := []int{1, 3, 5, 7, 4, 9, 10}

	in := make(chan int)
	out := make(chan int)

	wg.Add(2)

	go worker(wg, in, out)
	go printer(wg, out)

	for _, v := range n {
		in <- v
	}
	close(in)

	wg.Wait()
}