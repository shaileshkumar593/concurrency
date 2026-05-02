13. Timeout Pattern
Use case
Avoid waiting forever.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)

	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
Pros
Prevents blocking forever

Useful for external calls

Cons
time.After in loops can allocate repeatedly

Prefer time.NewTimer for hot loops