20. Rate Limiter Pattern
Use case
Limit requests per time window.

package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := time.Tick(500 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		<-limiter
		fmt.Println("Request allowed:", i, "at", time.Now())
	}
}
Pros
Simple

Useful for API throttling

Cons
time.Tick cannot be stopped

Prefer time.NewTicker in production