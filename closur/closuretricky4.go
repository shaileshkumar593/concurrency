package main

import (
	"fmt"
	"sync"
)

// A closure is a function that captures variables from its surrounding scope.

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}