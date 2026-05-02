23. Backpressure Pattern
Use case
Reject or slow down when system is overloaded.

package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int, 3)

	for i := 1; i <= 10; i++ {
		select {
		case queue <- i:
			fmt.Println("accepted:", i)

		default:
			fmt.Println("rejected due to backpressure:", i)
		}
	}

	close(queue)

	for job := range queue {
		fmt.Println("processing:", job)
		time.Sleep(time.Second)
	}
}
Pros
Protects system

Prevents memory explosion

Cons
Drops/rejects work

Requires client retry strategy