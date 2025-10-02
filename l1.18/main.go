package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
    Increment()
    Value() int64
    Add(delta int64)
}

type MutexCounter struct {
    mu    sync.RWMutex
    value int64
}

func NewMutexCounter() *MutexCounter {
    return &MutexCounter{}
}

func (c *MutexCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *MutexCounter) Value() int64 {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.value
}

func (c *MutexCounter) Add(delta int64) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value += delta
}

type AtomicCounter struct {
    value int64
}

func NewAtomicCounter() *AtomicCounter {
    return &AtomicCounter{}
}

func (c *AtomicCounter) Increment() {
    atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.value)
}

func (c *AtomicCounter) Add(delta int64) {
    atomic.AddInt64(&c.value, delta)
}

type ChannelCounter struct {
    increment chan int64
    value     chan int64
    stop      chan struct{}
}

func NewChannelCounter() *ChannelCounter {
    counter := &ChannelCounter{
        increment: make(chan int64),
        value:     make(chan int64),
        stop:      make(chan struct{}),
    }
    
    go counter.run()
    return counter
}

func (c *ChannelCounter) run() {
    var count int64
    for {
        select {
        case delta := <-c.increment:
            count += delta
        case c.value <- count:
        case <-c.stop:
            return
        }
    }
}

func (c *ChannelCounter) Increment() {
    c.increment <- 1
}

func (c *ChannelCounter) Value() int64 {
    return <-c.value
}

func (c *ChannelCounter) Add(delta int64) {
    c.increment <- delta
}

func (c *ChannelCounter) Stop() {
    close(c.stop)
}

func testCounter(counter Counter, goroutines int, incrementsPerGoroutine int) {
    var wg sync.WaitGroup
    start := time.Now()
    
    for i := 0; i < goroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < incrementsPerGoroutine; j++ {
                counter.Increment()
            }
        }()
    }
    
    wg.Wait()
    duration := time.Since(start)
    
    expected := int64(goroutines * incrementsPerGoroutine)
    actual := counter.Value()
    
    fmt.Printf("Goroutine: %d, increments on goroutine: %d\n", goroutines, incrementsPerGoroutine)
    fmt.Printf("Expected value: %d\n", expected)
    fmt.Printf("Actual value: %d\n", actual)
    fmt.Printf("Time: %v\n", duration)
    fmt.Printf("Corretness: %t\n\n", expected == actual)
}

func main() {
    fmt.Println("Testing Mutex Counter")
    mutexCounter := NewMutexCounter()
    testCounter(mutexCounter, 1000, 100)
    
    fmt.Println("Testing Atomic Counter")
    atomicCounter := NewAtomicCounter()
    testCounter(atomicCounter, 1000, 100)
    
    fmt.Println("Testing Channel Counter")
    channelCounter := NewChannelCounter()
    testCounter(channelCounter, 1000, 100)
}