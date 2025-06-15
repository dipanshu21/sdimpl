package main

import (
	"fmt"
	"time"
)

const (
	// Size defines the number of elements to insert
	Size = 1000000
)

// BenchmarkResult holds the timing results for a benchmark
type BenchmarkResult struct {
	name     string
	duration time.Duration
}

func main() {
	// Test different slice initialization strategies
	results := []BenchmarkResult{
		benchmarkEmptySlice(),
		benchmarkPreallocatedSlice(),
	}

	// Print results in a formatted table
	fmt.Println("\nBenchmark Results:")
	fmt.Println("----------------------------------------")
	fmt.Printf("%-30s | %s\n", "Strategy", "Duration")
	fmt.Println("----------------------------------------")
	for _, result := range results {
		fmt.Printf("%-30s | %v\n", result.name, result.duration)
	}
}

// measureTime executes a function and measures its execution time
func measureTime(f func(), name string) BenchmarkResult {
	start := time.Now()
	f()
	return BenchmarkResult{
		name:     name,
		duration: time.Since(start),
	}
}

// insertElements inserts n elements into a slice
// The slice is passed by value, so we need to return the new slice
func insertElements(slice []int, n int) []int {
	// Insert elements
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return slice
}

// benchmarkEmptySlice measures performance of appending to empty slice
func benchmarkEmptySlice() BenchmarkResult {
	var emptySlice []int
	return measureTime(
		func() {
			_ = insertElements(emptySlice, Size)
		},
		"Empty Slice",
	)
}

// benchmarkPreallocatedSlice measures performance with pre-allocated length
func benchmarkPreallocatedSlice() BenchmarkResult {
	preAllocated := make([]int, 0, Size)
	return measureTime(
		func() {
			_ = insertElements(preAllocated, Size)
		},
		"Preallocated Capacity",
	)
}
