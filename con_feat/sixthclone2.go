package main

import (
	"fmt"
	"sync"
)

func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		jobs <- id*10 + i
	}
}

func consumer(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result := job * 2
		fmt.Println("Consumer", id, "processed", job, "result:", result)
		results <- result
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup

	for i := 1; i <= 3; i++ {
		producerWG.Add(1)
		go producer(i, jobs, &producerWG)
	}

	for i := 1; i <= 2; i++ {
		consumerWG.Add(1)
		go consumer(i, jobs, results, &consumerWG)
	}

	go func() {
		producerWG.Wait()
		close(jobs)
	}()

	go func() {
		consumerWG.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Final result:", result)
	}
}