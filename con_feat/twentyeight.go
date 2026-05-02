28. Channel Ownership Pattern
Rule
The goroutine that sends should usually close the channel.

package main

import "fmt"

func generate() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 1; i <= 5; i++ {
			out <- i
		}
	}()

	return out
}

func main() {
	ch := generate()

	for val := range ch {
		fmt.Println(val)
	}
}
Pros
Avoids panic from sending on closed channel

Clear ownership

Cons
Needs discipline

Multiple senders require coordination