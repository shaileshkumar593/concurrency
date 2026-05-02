10. Pipeline Pattern
Use case
Multiple processing stages.

package main

import "fmt"

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n * 2
		}
	}()

	return out
}

func main() {
	nums := generate(1, 2, 3, 4)
	squares := square(nums)
	results := double(squares)

	for result := range results {
		fmt.Println(result)
	}
}
Pros
Clean data flow

Easy to compose stages

Cons
Goroutine leaks if consumer stops early

Error handling is harder