package main

import (
	"fmt"
	"sync"
	"time"
)

// Method 1: sync.Mutex
type SafeMapMutex struct {
	mu sync.Mutex
	m  map[int]string
}

func NewSafeMapMutex() *SafeMapMutex {
	return &SafeMapMutex{
		m: make(map[int]string),
	}
}

func (sm *SafeMapMutex) Set(key int, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMapMutex) Get(key int) (string, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMapMutex) Delete(key int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

func (sm *SafeMapMutex) Len() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return len(sm.m)
}

func mutexExample() {
	fmt.Println("\nMethod 1: sync.Mutex")
	
	safeMap := NewSafeMapMutex()
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			safeMap.Set(id, fmt.Sprintf("worker-%d", id))
			time.Sleep(10 * time.Millisecond)
		}(i)
	}
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if val, ok := safeMap.Get(id); ok {
				fmt.Printf("Reader: key=%d, value=%s\n", id, val)
			}
			time.Sleep(10 * time.Millisecond)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Total elements: %d\n", safeMap.Len())
}

// Method 2: sync.RWMutex
type SafeMapRWMutex struct {
	mu sync.RWMutex
	m  map[int]string
}

func NewSafeMapRWMutex() *SafeMapRWMutex {
	return &SafeMapRWMutex{
		m: make(map[int]string),
	}
}

func (sm *SafeMapRWMutex) Set(key int, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMapRWMutex) Get(key int) (string, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMapRWMutex) Delete(key int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

func (sm *SafeMapRWMutex) Keys() []int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	keys := make([]int, 0, len(sm.m))
	for k := range sm.m {
		keys = append(keys, k)
	}
	return keys
}

func rwMutexExample() {
	fmt.Println("\nMethod 2: sync.RWMutex")
	
	safeMap := NewSafeMapRWMutex()
	var wg sync.WaitGroup
	
	// 5 писателей
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			safeMap.Set(id, fmt.Sprintf("data-%d", id))
			fmt.Printf("Get: key=%d\n", id)
		}(i)
	}
	
	time.Sleep(50 * time.Millisecond)
	
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := id % 5
			if val, ok := safeMap.Get(key); ok {
				fmt.Printf("Reader %d: key=%d, value=%s\n", id, key, val)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Keys in map: %v\n", safeMap.Keys())
}

// Method 3: sync.Map
func syncMapExample() {
	fmt.Println("\nMethod 3: sync.Map")
	
	var sm sync.Map
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Store(id, fmt.Sprintf("sync-value-%d", id))
			fmt.Printf("Save: %d\n", id)
		}(i)
	}
	
	wg.Wait()
	
	fmt.Println("\nRead all elements:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("  key=%v, value=%v\n", key, value)
		return true
	})
	
	sm.Delete(5)
	
	if val, ok := sm.Load(5); ok {
		fmt.Printf("Find: %v\n", val)
	} else {
		fmt.Println("Key 5 has been removed")
	}
	
	actual, loaded := sm.LoadOrStore(100, "new value")
	if loaded {
		fmt.Printf("Value already existed: %v\n", actual)
	} else {
		fmt.Printf("Value created: %v\n", actual)
	}
}


func main() {
	fmt.Println("Demonstrate all method's")
	
	mutexExample()
	time.Sleep(200 * time.Millisecond)
	
	rwMutexExample()
	time.Sleep(200 * time.Millisecond)
	
	syncMapExample()
	time.Sleep(200 * time.Millisecond)

	
	fmt.Println("\nDemonstration have ended")
}