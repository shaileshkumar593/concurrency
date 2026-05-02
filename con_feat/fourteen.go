14. Context Timeout Pattern
Use case
Cancel work after deadline.

package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("work done")
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := doWork(ctx)
	if err != nil {
		fmt.Println("error:", err)
	}
}
Pros
Standard production pattern

Works across API boundaries

Cons
Functions must accept context

Cancellation is cooperative