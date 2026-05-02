// 7. Worker Pool Pattern
// Use case
// -----Limit concurrency.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
		time.Sleep(time.Millisecond * 500)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Result:", result)
	}
}

/*
Pros
Prevents unlimited goroutines

Good for APIs, batch jobs, queues

Cons
Need tuning worker count

Too few workers = slow

Too many workers = resource pressure
*/
