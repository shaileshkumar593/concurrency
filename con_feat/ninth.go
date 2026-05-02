9. Fan-In Pattern
Use case
Merge multiple channels into one.

package main

import (
	"fmt"
	"sync"
)

func producer(name string, out chan<- string) {
	for i := 1; i <= 3; i++ {
		out <- fmt.Sprintf("%s-%d", name, i)
	}
	close(out)
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)

		go func(c <-chan string) {
			defer wg.Done()

			for msg := range c {
				out <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer("A", ch1)
	go producer("B", ch2)

	merged := fanIn(ch1, ch2)

	for msg := range merged {
		fmt.Println(msg)
	}
}
Pros
Combines multiple sources

Useful in pipelines

Cons
Output ordering is nondeterministic

Need careful closing
