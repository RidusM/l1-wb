package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func method1ExitByFlag() {
	fmt.Println("\nMethod 1: exit on condition (flag)")
	
	shouldStop := false
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			if shouldStop {
				fmt.Println("Gouroutine stopped on flag")
				done <- true
				return
			}
			fmt.Printf("Work... %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	time.Sleep(1 * time.Second)
	shouldStop = true
	<-done
}

func method2AtomicFlag() {
	fmt.Println("\nMethod 2: atomic flag")
	
	var stopFlag atomic.Bool
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			if stopFlag.Load() {
				fmt.Println("Gouroutine stopped on atomic flag")
				done <- true
				return
			}
			fmt.Printf("Atomic work... %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	time.Sleep(1 * time.Second)
	stopFlag.Store(true)
	<-done
}

func method3CloseChannel() {
	fmt.Println("\nMethod 3: close channel")
	
	stop := make(chan struct{})
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			select {
			case <-stop:
				fmt.Println("Stop signal recieved via channel")
				done <- true
				return
			default:
				fmt.Printf("Work with channel... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	close(stop)
	<-done
}

func method4ChannelMessage() {
	fmt.Println("\nMethod 4: channel with message")
	
	stop := make(chan bool)
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			select {
			case <-stop:
				fmt.Println("Stop signal recieved via channel")
				done <- true
				return
			default:
				fmt.Printf("Processing... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	stop <- true
	<-done
}

func method5Context() {
	fmt.Println("\nMethod 5: Context")
	
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Context canceled: %v\n", ctx.Err())
				done <- true
				return
			default:
				fmt.Printf("Work with context... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	cancel()
	<-done
}

func method6ContextTimeout() {
	fmt.Println("\nMethod 6: Context with timeout")
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Timeout expired: %v\n", ctx.Err())
				done <- true
				return
			default:
				fmt.Printf("Work until timeout... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	<-done
}

func method7ContextDeadline() {
	fmt.Println("\nMethod 7: Context with Deadline")
	
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	done := make(chan bool)
	
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Deadline achieved: %v\n", ctx.Err())
				done <- true
				return
			default:
				fmt.Printf("Work until deadline... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	<-done
}

func method8Goexit() {
	fmt.Println("\nMethod 8: runtime.Goexit()")
	
	stop := make(chan bool)
	done := make(chan bool)
	
	go func() {
		defer func() {
			fmt.Println("Defer before exit")
			done <- true
		}()
		
		for i := 0; ; i++ {
			select {
			case <-stop:
				fmt.Println("Call runtime.Goexit()")
				runtime.Goexit()
			default:
				fmt.Printf("Work with Goexit... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	stop <- true
	<-done
}

func method9WaitGroup() {
	fmt.Println("\nMethod 9: WaitGroup with channel")
	
	var wg sync.WaitGroup
	stop := make(chan struct{})
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-stop:
				fmt.Println("Stop in WaitGroup")
				return
			default:
				fmt.Printf("WaitGroup work... %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	close(stop)
	wg.Wait()
	fmt.Println("All goroutines completed")
}

func method10MultipleGoroutines() {
	fmt.Println("\nMethod 10: multiple goroutines with one channel ===")
	
	stop := make(chan struct{})
	var wg sync.WaitGroup
	
	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for i := 0; ; i++ {
				select {
				case <-stop:
					fmt.Printf("Worker %d stopped\n", workerID)
					return
				default:
					fmt.Printf("Worker %d work... %d\n", workerID, i)
					time.Sleep(300 * time.Millisecond)
				}
			}
		}(id)
	}
	
	time.Sleep(1 * time.Second)
	close(stop)
	wg.Wait()
	fmt.Println("All worker's stopped")
}

func method11Timer() {
	fmt.Println("\nMethod 11: Stop via ticker")
	
	ticker := time.NewTicker(200 * time.Millisecond)
	stop := make(chan bool)
	done := make(chan bool)
	
	go func() {
		i := 0
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Tick %d\n", i)
				i++
			case <-stop:
				ticker.Stop()
				fmt.Println("Ticker stopped")
				done <- true
				return
			}
		}
	}()
	
	time.Sleep(1 * time.Second)
	stop <- true
	<-done
}

func method12Panic() {
	fmt.Println("\nMethod 12: Panic")
	
	done := make(chan bool)
	
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
				done <- true
			}
		}()
		
		for i := 0; i < 5; i++ {
			fmt.Printf("Work before panic... %d\n", i)
			time.Sleep(200 * time.Millisecond)
			if i == 3 {
				panic("force stop")
			}
		}
	}()
	
	<-done
}

func main() {
	fmt.Println("Demonstrate all method's")
	
	method1ExitByFlag()
	time.Sleep(500 * time.Millisecond)
	
	method2AtomicFlag()
	time.Sleep(500 * time.Millisecond)
	
	method3CloseChannel()
	time.Sleep(500 * time.Millisecond)
	
	method4ChannelMessage()
	time.Sleep(500 * time.Millisecond)
	
	method5Context()
	time.Sleep(500 * time.Millisecond)
	
	method6ContextTimeout()
	time.Sleep(500 * time.Millisecond)
	
	method7ContextDeadline()
	time.Sleep(500 * time.Millisecond)
	
	method8Goexit()
	time.Sleep(500 * time.Millisecond)
	
	method9WaitGroup()
	time.Sleep(500 * time.Millisecond)
	
	method10MultipleGoroutines()
	time.Sleep(500 * time.Millisecond)
	
	method11Timer()
	time.Sleep(500 * time.Millisecond)
	
	method12Panic()
	
	fmt.Println("\nDemonstrate have ended")
}