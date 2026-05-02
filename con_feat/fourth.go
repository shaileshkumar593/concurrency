//  4. Buffered Channel Pattern
/*
Use case
Allow limited buffering between sender and receiver.*/


package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
Pros
	Reduces blocking

	Useful for queues

Cons
	Buffer can fill

	Can hide slow consumer problem
*/
