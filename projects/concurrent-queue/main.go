package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
}

func (cq *ConcurrentQueue) Enqueue(value int32) {
	cq.queue = append(cq.queue, value)
}

func (cq *ConcurrentQueue) Dequeue() int32 {
	if len(cq.queue) == 0 {
		panic("Queue is empty, cannot dequeue")
	}

	item := cq.queue[0]
	cq.queue = cq.queue[1:]

	return item
}

func (cq *ConcurrentQueue) Size() int {
	return len(cq.queue)
}

func runQueueV1() {
	q1 := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	fmt.Println("Queue size after enqueuing 3 items:", q1.Size())
	fmt.Println("Dequeue item:", q1.Dequeue())
	fmt.Println("Dequeue item:", q1.Dequeue())
	fmt.Println("Dequeue item:", q1.Dequeue())
	fmt.Println("Dequeue item:", q1.Dequeue())
}

func runQueueV2() {
	var wg sync.WaitGroup
	q1 := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q1.Enqueue(rand.Int31())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Queue size after enqueuing 1,000,000 items:", q1.Size())
}

type ConcurrentQueueV3 struct {
	queue []int32
	mu    sync.Mutex
}

func (cq *ConcurrentQueueV3) Enqueue(value int32) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.queue = append(cq.queue, value)
}

func (cq *ConcurrentQueueV3) Dequeue() int32 {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	if len(cq.queue) == 0 {
		panic("Queue is empty, cannot dequeue")
	}
	item := cq.queue[0]
	cq.queue = cq.queue[1:]
	return item
}

func (cq *ConcurrentQueueV3) Size() int {
	return len(cq.queue)
}

func runQueueV3() {
	var wgEnqueue, wgDequeue sync.WaitGroup
	q1 := ConcurrentQueueV3{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1000000; i++ {
		wgEnqueue.Add(1)
		go func() {
			q1.Enqueue(rand.Int31())
			wgEnqueue.Done()
		}()
	}

	for i := 0; i < 1000000; i++ {
		wgDequeue.Add(1)
		go func() {
			q1.Dequeue()
			wgDequeue.Done()
		}()
	}

	wgEnqueue.Wait()
	wgDequeue.Wait()
	fmt.Println("Queue size after enqueuing 1,000,000 items:", q1.Size())
}

func main() {
	//runQueueV1()
	//runQueueV2()
	runQueueV3()
}
