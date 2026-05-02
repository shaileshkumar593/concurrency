27. Context Tree Pattern
Use case
Cancel child operations from parent.

package main

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "cancelled")
			return

		default:
			fmt.Println(name, "working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	parentCtx, cancelParent := context.WithCancel(context.Background())

	childCtx1, _ := context.WithCancel(parentCtx)
	childCtx2, _ := context.WithCancel(parentCtx)

	go child(childCtx1, "child-1")
	go child(childCtx2, "child-2")

	time.Sleep(2 * time.Second)
	cancelParent()

	time.Sleep(time.Second)
}
Pros
Clean cancellation propagation

Important for APIs/microservices

Cons
Cancellation does not kill goroutine automatically

Code must check ctx.Done()