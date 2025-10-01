package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := 5 * time.Second

	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
	}()

	done := time.After(timeout)
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				fmt.Println("Channel close")
				return
			}
			fmt.Println("Recieved value:", value)
		case <-done:
			fmt.Println("Timeout expired, terminating...")
			return
		}
	}
}