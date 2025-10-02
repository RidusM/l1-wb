package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func method1_ContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker1(ctx, &wg, i)
	}
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("Signal SIGINT/SIGTERM received")
	
	cancel()
	wg.Wait()

	fmt.Println("Gracefull shutdown")
}

func worker1(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: SIGINT/SIGTERM received\n", id)
			return
		default:
			fmt.Printf("Worker %d: do something...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func method2_DoneChannel() {
	done := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker2(done, &wg, i)
	}
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("Signal SIGINT/SIGTERM received")
	
	close(done)
	wg.Wait()

	fmt.Println("Gracefull shutdown")
}

func worker2(done <-chan struct{}, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	
	for {
		select {
		case <-done:
			fmt.Printf("Worker %d: SIGINT/SIGTERM received\n", id)
			return
		default:
			fmt.Printf("Worker %d: do something...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func method3_ContextWithTimeout() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker3(ctx, &wg, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("Signal SIGINT/SIGTERM received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	
	cancel()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("Gracefull shutdown")
	case <-shutdownCtx.Done():
		fmt.Println("Timeout! force exit")
	}
}

func worker3(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: start shutdown\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d: cleanup complete\n", id)
			return
		default:
			fmt.Printf("Worker %d: do something...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func method4_BroadcastChannel() {
	quit := make(chan struct{})
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker4(quit, &wg, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("Signal SIGINT/SIGTERM received")
	
	close(quit)
	wg.Wait()
	fmt.Println("Gracefull shutdown")
}

func worker4(quit <-chan struct{}, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-quit:
			fmt.Printf("Worker %d: broadcast signal received\n", id)
			return
		case <-ticker.C:
			fmt.Printf("Worker %d: do something...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func method5_Errgroup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker5(ctx, &wg, i, cancel)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("Signal SIGINT/SIGTERM received")
	
	cancel()
	wg.Wait()
	
	fmt.Println("Gracefull shutdown")
}

func worker5(ctx context.Context, wg *sync.WaitGroup, id int, cancel context.CancelFunc) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: context done\n", id)
			return
		default:
			fmt.Printf("Worker %d: do something...\n", id)
			time.Sleep(500 * time.Millisecond)

			if id == 2 {
				cancel()
			}
		}
	}
}

func method6_AtomicFlag() {
	var (
		wg       sync.WaitGroup
		shutdown int32 
	)
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker6(&shutdown, &wg, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("Signal SIGINT/SIGTERM received")
	
	atomic.StoreInt32(&shutdown, 1)
	
	wg.Wait()
	fmt.Println("Gracefull shutdown")
}

func worker6(shutdown *int32, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	
	for {
		if atomic.LoadInt32(shutdown) == 1{
			fmt.Printf("Worker %d: set exited flag\n", id)
			return
		}
		
		fmt.Printf("Worker %d: do something...\n", id)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Uncomment one of:
	method1_ContextCancel() // RECOMMEND, idiomatic, support timeout's/deadline's
	// method2_DoneChannel() // SIMPLE PROGS, min dependencies
	// method3_ContextWithTimeout() // PRODUCTION, guarantee completion in time
	// method4_BroadcastChannel() // MANY GOROUTINES, don't require worker loop, notification via close()
	// method5_Errgroup() // COMPLEX LOGIC, auto err handling
	// method6_AtomicFlag() // NOT RECOMMEND, min overhead, BUT less secure, risk memory ordering
}