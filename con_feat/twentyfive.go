25. Once Pattern
Use case
Run initialization only once.

package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialized only once")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}

	wg.Wait()
}
Pros
Thread-safe initialization

Great for singletons/config loading

Cons
Cannot reset easily

If init fails, retry logic is not built-in
