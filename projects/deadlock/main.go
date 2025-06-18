package main

import (
	"sync"
	"time"
)

const (
	NumOfThreads = 6 // Number of threads to run concurrently
)

/*
*
  - This function simulates a deadlock scenario where multiple threads try to acquire locks on multiple mutexes.

Sample output: first number if the mutex id, and subsequent numbers are the thread ids that are waiting/holding the lock on the mutex

0 Locked (5) Waiting(3) Waiting(0) Waiting(2)
1 Locked (4) Waiting(3) Waiting(1) Waiting(0)
2 Locked (5) Waiting(4) Waiting(1) Waiting(2)

0 Locked(3) Waiting(0) Waiting(2)
1 Locked (4) Waiting(3) Waiting(1) Waiting(0)
2 Locked(4) Waiting(1) Waiting(2)

0 Locked(3) Waiting(0) Waiting(2)
1 Locked(1) Waiting(3)  Waiting(0)
2 Locked(2) Waiting(1)
*/
func sampleDeadlock() {
	var mutexs [3]sync.Mutex // Mutex for synchronizing access to the queue
	var wg sync.WaitGroup
	wg.Add(NumOfThreads) // Add the number of threads to the wait group
	for i := 0; i < NumOfThreads; i++ {
		go func(i int) {
			// Simulate some work
			println("Thread", i, "is starting")
			println("Thread", i, "is waiting to lock mutex", i%3, "and mutex", (i+1)%3)
			mutexs[i%3].Lock() // Lock the mutex for the current thread
			println("Thread", i, "acquired lock on mutex", i%3)
			mutexs[(i+1)%3].Lock() // Lock the mutex for the current thread
			println("Thread", i, "acquired lock on mutex", (i+1)%3)
			defer mutexs[i%3].Unlock()     // Ensure the mutex is unlocked when done
			defer mutexs[(i+1)%3].Unlock() // Ensure the mutex is unlocked when done
			println("Thread", i, "is processing with locks held")
			time.Sleep(2 * time.Second) // Simulate some processing time
			// Simulate processing
			println("Thread", i, "is done processing releasing locks")
			wg.Done() // Signal that the thread is done
		}(i)
	}

	wg.Wait() // Wait for all threads to finish
	println("All threads have finished processing")
}

func main() {
	sampleDeadlock()
}
