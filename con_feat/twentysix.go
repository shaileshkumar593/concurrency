26. Condition Variable Pattern
Use case
Wait until condition becomes true.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	ready := false

	go func() {
		time.Sleep(time.Second)

		mu.Lock()
		ready = true
		cond.Signal()
		mu.Unlock()
	}()

	mu.Lock()
	for !ready {
		cond.Wait()
	}
	mu.Unlock()

	fmt.Println("Condition met")
}
Pros
Efficient waiting

Useful for advanced synchronization

Cons
Harder to understand

Easy to misuse

Channels are often simpler