29. Non-blocking Channel Send/Receive
Use case
Try operation without blocking.

package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 10:
		fmt.Println("sent")
	default:
		fmt.Println("channel full")
	}

	select {
	case val := <-ch:
		fmt.Println("received:", val)
	default:
		fmt.Println("nothing to receive")
	}
}
Pros
Avoids blocking

Useful for backpressure

Cons
Can drop work

Requires careful business logic