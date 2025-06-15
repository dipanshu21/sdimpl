package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const (
	// Size defines the number of elements to insert
	Size      = 1000 // Number of DB calls to simulate
	MAX_DELAY = 500  // Maximum delay in milliseconds for each DB call
)

var m = sync.Mutex{}      // Mutex to protect shared data
var rwm = sync.RWMutex{}  // Read-Write Mutex to protect shared data
var wg = sync.WaitGroup{} // WaitGroup to wait for all goroutines to finish
var dbData = generateRandomDBData(Size)
var timeDelays = generateRandomDelays(Size) // Random delays for each DB call
var results = []string{}                    // Results of DB calls

func generateRandomDBData(n int) []string {
	// Simulate generating random data for the database
	// In a real application, this would fetch data from a database or an API
	var randomData []string = make([]string, 0, n)

	for i := 0; i < n; i++ {
		randomData = append(randomData, fmt.Sprintf("data%d", i+1)) // Simulate random data
	}

	//fmt.Println("Generated random data for DB calls:", randomData)

	return randomData
}

func generateRandomDelays(n int) []float32 {
	// Simulate generating random data for the database
	// In a real application, this would fetch data from a database or an API
	var randomDelays []float32 = make([]float32, 0, n)

	for i := 0; i < n; i++ {
		randomDelays = append(randomDelays, generateRandomDelay())
	}

	//fmt.Println("Generated random delays for DB calls:", randomDelays)

	return randomDelays
}

func generateRandomDelay() float32 {
	// Simulate generating a random delay for each DB call
	var randomDelay float32 = rand.Float32() * MAX_DELAY // Random delay between 0 and 1000 milliseconds
	return randomDelay
}

func main() {

	// fmt.Println("Now executing DB calls without goroutines...")
	// executeDBCallsWithoutGoroutines()

	// fmt.Println("\n\nNow executing DB calls with goroutines...")
	// executeDBCallsWithGoroutines_Channels()

	// fmt.Println("\n\nNow executing DB calls with goroutines and WaitGroups...")
	// executeDBCallsWithGoroutines_WaitGroups()

	// fmt.Println("\n\nAll DB calls modifying same data bad implementation...")
	// executeDBCallsWithGoroutines_ModifyingSameData_Bad()
	// fmt.Println("\n\nResults of DB calls:", results)
	// fmt.Println("\n\nResults of DB calls:", len(results))

	// fmt.Println("\n\nAll DB calls modifying same data good implementation...")
	// executeDBCallsWithGoroutines_ModifyingSameData_Good()
	// fmt.Println("\n\nResults of DB calls:", results)

	fmt.Println("\n\nAll DB calls modifying same data even better implementation...")
	executeDBCallsWithGoroutines_ModifyingSameData_EvenBetter()
	fmt.Println("\n\nResults of DB calls:", results)
	fmt.Println("\n\nResults of DB calls:", len(results))

	//rwm.Lock() // Lock the Read-Write Mutex to read the results
	// fmt.Println("\n\nResults of DB calls:", results)
	// fmt.Println("\n\nResults of DB calls:", len(results))
	//rwm.Unlock() // Unlock the Read-Write Mutex after reading the results

}

func executeDBCallsWithoutGoroutines() {
	tStart := time.Now()

	for i := 0; i < len(dbData); i++ {
		dbCall(i, true) // Call the DB function directly
	}

	fmt.Printf("\nAll DB calls completed in %v seconds\n", time.Since(tStart).Seconds())
}

/**
 * This function executes DB calls using goroutines and waits for all of them to complete using channels.
 * It creates a goroutine for each DB call and uses a channel to signal when each goroutine is done.
 * The main function waits for all goroutines to finish before proceeding.
 * By waiting on same number of signals as the number of goroutines, we ensure that all DB calls are completed before exiting the function.
 */
func executeDBCallsWithGoroutines_Channels() {
	tStart := time.Now()

	// Create a channel to synchronize goroutines
	done := make(chan bool)

	for i := 0; i < len(dbData); i++ {
		go func(i int) {
			dbCall(i, true) // Call the DB function in a goroutine
			done <- true    // Signal that the goroutine is done
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < len(dbData); i++ {
		<-done
	}

	fmt.Printf("\nAll DB calls completed in %v seconds\n", time.Since(tStart).Seconds())
}

/**
 * This function executes DB calls using goroutines and waits for all of them to complete using WaitGroups.
 * The Wait group works very simply by maintaining a counter that tracks the number of goroutines that are running.
 */
func executeDBCallsWithGoroutines_WaitGroups() {
	tStart := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(i int) {
			dbCall(i, true) // Call the DB function in a goroutine
			wg.Done()       // Decrement the WaitGroup counter when done
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("\nAll DB calls completed in %v seconds\n", time.Since(tStart).Seconds())
}

/**
 * This function executes DB calls using goroutines and waits for all of them to complete using WaitGroups.
 * The Wait group works very simply by maintaining a counter that tracks the number of goroutines that are running.
 * However, this implementation is bad because it modifies the same data (results slice) from multiple goroutines concurrently,
 * which can lead to a situation where the results are not consistent or may even cause a panic due to concurrent writes.
 */
func executeDBCallsWithGoroutines_ModifyingSameData_Bad() {
	tStart := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(i int) {
			dbCall(i, false)                    // Call the DB function in a goroutine
			dbResult := dbData[i]               // Simulate fetching data from the database
			results = append(results, dbResult) // Store the result in the results slice
			wg.Done()                           // Decrement the WaitGroup counter when done
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("\nAll DB calls completed in %v seconds\n", time.Since(tStart).Seconds())
}

/**
 * This function executes DB calls using goroutines and waits for all of them to complete using WaitGroups.
 * It also uses a mutex to protect shared data (results slice) from concurrent writes.
 * This implementation is good because it ensures that only one goroutine can modify the results slice at a time,
 * When m.Lock() is called, it locks the mutex, preventing other goroutines from modifying the results slice until the lock is released with m.Unlock().
 */
func executeDBCallsWithGoroutines_ModifyingSameData_Good() {
	tStart := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(i int) {
			dbCall(i, false)                    // Call the DB function in a goroutine
			dbResult := dbData[i]               // Simulate fetching data from the database
			m.Lock()                            // Lock the mutex to protect shared data
			results = append(results, dbResult) // Store the result in the results slice
			m.Unlock()                          // Unlock the mutex after modifying shared data
			wg.Done()                           // Decrement the WaitGroup counter when done
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("\nAll DB calls completed in %v seconds\n", time.Since(tStart).Seconds())
}

/**
 *
 */
func executeDBCallsWithGoroutines_ModifyingSameData_EvenBetter() {
	tStart := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func(i int) {
			dbCall(i, false)                    // Call the DB function in a goroutine
			dbResult := dbData[i]               // Simulate fetching data from the database
			rwm.Lock()                          // Lock the mutex to protect shared data
			results = append(results, dbResult) // Store the result in the results slice
			rwm.Unlock()                        // Unlock the mutex after modifying shared data
			wg.Done()                           // Decrement the WaitGroup counter when done
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("\nAll DB calls completed in %v seconds\n", time.Since(tStart).Seconds())
}

func dbCall(i int, print bool) {
	delay := timeDelays[i]
	time.Sleep(time.Duration(delay) * time.Millisecond)

	if print {
		dbResult := dbData[i] // Simulate fetching data from the database
		fmt.Printf("\nDB Call Completed after %v milliseconds and the response is: %v\n", delay, dbResult)
	}
}
