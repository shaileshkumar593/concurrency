// 2. WaitGroup Pattern
/*Use case
Wait for multiple goroutines to complete.*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker started:", id)
	time.Sleep(time.Second)
	fmt.Println("Worker finished:", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}

/*
Pros
	Simple

	Best for waiting on fixed goroutines

Cons
	No built-in error handling

	No cancellation
	Must call Done() correctly
*/




















































