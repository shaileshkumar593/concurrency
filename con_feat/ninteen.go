19. Error Group Pattern
Use case
Run goroutines and stop on first error.

package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		time.Sleep(time.Second)
		fmt.Println("task 1 done")
		return nil
	})

	g.Go(func() error {
		time.Sleep(500 * time.Millisecond)
		return errors.New("task 2 failed")
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("task 3 cancelled")
			return ctx.Err()
		case <-time.After(2 * time.Second):
			fmt.Println("task 3 done")
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("error:", err)
	}
}
Install:

go get golang.org/x/sync/errgroup
Pros
Best production pattern for parallel tasks

Handles error + cancellation

Cons
External package

Stops on first error by default