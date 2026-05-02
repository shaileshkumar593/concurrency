//Basic Goroutine
//Use case
//Run work concurrently.

package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Println("Worker started:", id)
	time.Sleep(time.Second)
	fmt.Println("Worker finished:", id)
}

func main() {
	go worker(1)
	go worker(2)

	time.Sleep(2 * time.Second)
}
/*
Pros
	Very lightweight

	Easy to start concurrent work

Cons
	main() may exit before goroutines finish

	No error handling

	No cancellation
*/
