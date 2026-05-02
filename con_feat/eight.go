8. Fan-Out Pattern
Use case
Distribute work to many workers.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "got job", job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}
Pros
Parallel processing

Easy to scale consumers

Cons
Ordering not guaranteed

Load depends on worker speed