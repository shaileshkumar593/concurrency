22. Token Bucket Pattern
Use case
Allow burst + refill over time.

package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	tokens chan struct{}
}

func NewTokenBucket(capacity int, refillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		tokens: make(chan struct{}, capacity),
	}

	for i := 0; i < capacity; i++ {
		tb.tokens <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(refillInterval)
		defer ticker.Stop()

		for range ticker.C {
			select {
			case tb.tokens <- struct{}{}:
			default:
			}
		}
	}()

	return tb
}

func (tb *TokenBucket) Allow() bool {
	select {
	case <-tb.tokens:
		return true
	default:
		return false
	}
}

func main() {
	limiter := NewTokenBucket(3, time.Second)

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Println("allowed:", i)
		} else {
			fmt.Println("blocked:", i)
		}

		time.Sleep(300 * time.Millisecond)
	}
}
Pros
Supports burst

Common in API gateways

Cons
Needs background goroutine

More complex than fixed rate
