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

	"golang.org/x/sync/errgroup"
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
	fmt.Println("\nSignal SIGINT/SIGTERM received")

	cancel()
	wg.Wait()

	fmt.Println("Graceful shutdown")
}

func worker1(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: SIGINT/SIGTERM received, exiting\n", id)
			return
		case <-ticker.C:
			fmt.Printf("Worker %d: do something...\n", id)
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
	fmt.Println("\nSignal SIGINT/SIGTERM received")

	close(done)
	wg.Wait()

	fmt.Println("Graceful shutdown")
}

func worker2(done <-chan struct{}, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Printf("Worker %d: SIGINT/SIGTERM received, exiting\n", id)
			return
		case <-ticker.C:
			fmt.Printf("Worker %d: do something...\n", id)
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
	fmt.Println("\nSignal SIGINT/SIGTERM received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	cancel()

	waitChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitChan)
	}()

	select {
	case <-waitChan:
		fmt.Println("Graceful shutdown complete")
	case <-shutdownCtx.Done():
		fmt.Println("Timeout! Force exit, some workers might be stuck")
	}
}

func worker3(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: start cleanup\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d: cleanup complete\n", id)
			return
		case <-ticker.C:
			fmt.Printf("Worker %d: do something...\n", id)
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
	fmt.Println("\nSignal SIGINT/SIGTERM received")

	close(quit)
	wg.Wait()
	fmt.Println("Graceful shutdown")
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
		}
	}
}

func method5_Errgroup() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for i := 1; i <= 3; i++ {
		id := i
		g.Go(func() error {
			return worker5(ctx, id)
		})
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		fmt.Println("\nSignal SIGINT/SIGTERM received")
		cancel()
	}()

	if err := g.Wait(); err != nil {
		fmt.Printf("Errgroup finished with error: %v\n", err)
	} else {
		fmt.Println("Graceful shutdown via Errgroup complete")
	}
}

func worker5(ctx context.Context, id int) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: context done\n", id)
			return ctx.Err()
		case <-ticker.C:
			fmt.Printf("Worker %d: do something...\n", id)

			// demo fail: if worker 2 failed, 
			// errgroup cancel context for workers 1 and 3
			if id == 2 {
				// Uncomment for check auto-cancel
				// return fmt.Errorf("worker 2 failed encounterly")
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
	fmt.Println("\nSignal SIGINT/SIGTERM received")

	atomic.StoreInt32(&shutdown, 1)

	wg.Wait()
	fmt.Println("Graceful shutdown")
}
func worker6(shutdown *int32, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		if atomic.LoadInt32(shutdown) == 1 {
			fmt.Printf("Worker %d: detected shutdown flag via atomic\n", id)
			return
		}
		fmt.Printf("Worker %d: do something...\n", id)
	}
}

func main() {
	// method1_ContextCancel()
	// method2_DoneChannel()
	// method3_ContextWithTimeout()
	// method4_BroadcastChannel()
	// method5_Errgroup()
	// method6_AtomicFlag()
}