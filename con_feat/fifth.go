// 5. Producer Consumer Pattern
/*
Use case
Producer generates jobs, consumer processes them.*/


package main

import (
	"fmt"
	"sync"
)

func producer(jobs chan<- int) {
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Processing job:", job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go consumer(jobs, &wg)

	producer(jobs)

	wg.Wait()
}

/*
Pros
	Clean separation of producer and consumer

	Channel close signals completion

Cons
	Slow consumer can block producer

	Need proper close ownership
*/
