package main

import (
	"fmt"
	"sync"
)

func fanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for val := range ch1 {
			out <- val
		}
	}()

	go func() {
		defer wg.Done()
		for val := range ch2 {
			out <- val
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i <= 12; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	merged := fanIn(ch1, ch2)

	for val := range merged {
		fmt.Println(val)
	}
}