package main

import (
	"sync"
	"time"
)

var count int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

const ( // Number of goroutines to run concurrently
	numGoroutines = 1000000
)

func doCount_not_thread_safe() {
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrCount_not_thread_safe()
	}
}

func incrCount_not_thread_safe() {
	count++
	wg.Done()
}

func doCount_thread_safe() {
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrCount_thread_safe()
	}
}

func incrCount_thread_safe() {
	mutex.Lock()
	count++
	mutex.Unlock()
	wg.Done()
}

func main() {
	t0 := time.Now()
	count = 0
	doCount_not_thread_safe()
	wg.Wait()
	t1 := time.Now()
	duration := t1.Sub(t0)
	println("Duration not thread safe:", duration.Microseconds(), " ms")
	println("Final count:", count)

	t0 = time.Now()
	count = 0
	doCount_thread_safe()
	wg.Wait()
	t1 = time.Now()
	duration = t1.Sub(t0)
	println("Duration thread safe:", duration.Microseconds(), " ms")
	println("Final count:", count)
}
