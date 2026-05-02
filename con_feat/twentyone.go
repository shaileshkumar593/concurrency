21. Production Ticker Rate Limiter
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := 1; i <= 5; i++ {
		<-ticker.C
		fmt.Println("Request allowed:", i)
	}
}
Pros
Can stop ticker

Safer than time.Tick

Cons
Fixed rate only

No burst support