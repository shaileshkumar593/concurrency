15. Semaphore Pattern
Use case
Limit number of concurrent goroutines.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() {
				<-sem
			}()

			fmt.Println("Processing:", id)
			time.Sleep(time.Second)
		}(i)
	}

	wg.Wait()
}
Pros
Simple concurrency limit

Great for DB/API limits

Cons
Must release semaphore

No queue priority