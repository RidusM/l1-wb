package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	wFlags := flag.Int("workers", 10, "number of workers")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGTERM)
	defer cancel()

	jobs := make(chan int, *wFlags)
	var wg sync.WaitGroup

	for i := 0; i < *wFlags; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		defer close(jobs)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				jobs <- rand.Int()
			}
		}
	}()

	wg.Wait()

	fmt.Println("\nGraceful Shutdown")
}

func worker(ctx context.Context, counter int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker:%d, Result:%d\n", counter+1, job)
		case <-ctx.Done():
			return
		}
	}
}