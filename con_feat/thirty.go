30. Worker with Context + Jobs Channel
Production pattern
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker cancelled:", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Println("jobs closed, worker exit:", id)
				return
			}

			fmt.Println("worker", id, "processing job", job)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
	fmt.Println("done")
}
Pros
Safe shutdown

Prevents goroutine leaks

Good production template

Cons
More boilerplate

Need proper close/cancel order