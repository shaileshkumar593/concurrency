11. Pipeline with Cancellation
Use case
Stop all stages safely.

package main

import (
	"context"
	"fmt"
)

func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()

	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return

			case n, ok := <-in:
				if !ok {
					return
				}

				select {
				case <-ctx.Done():
					return
				case out <- n * n:
				}
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := generate(ctx, 1, 2, 3, 4, 5)
	results := square(ctx, nums)

	for result := range results {
		fmt.Println(result)

		if result == 9 {
			cancel()
			break
		}
	}
}
Pros
Prevents goroutine leaks

Production friendly

Cons
More code

Every stage must respect context