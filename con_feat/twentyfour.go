24. Graceful Shutdown Pattern
Use case
Stop workers cleanly.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", id)
			return

		default:
			fmt.Println("Worker running:", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
	fmt.Println("All workers stopped")
}
Pros
Avoids goroutine leaks

Production required

Cons
Every goroutine must respect context

Blocking calls may delay shutdown
