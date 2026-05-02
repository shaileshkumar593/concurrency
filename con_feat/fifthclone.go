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

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Consumer", id, "processing:", job)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	// Start multiple consumers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, jobs, &wg)
	}

	go producer(jobs)

	wg.Wait()
}