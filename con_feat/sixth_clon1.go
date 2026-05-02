package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producer(ctx context.Context, id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		job := id*10 + i

		select {
		case <-ctx.Done():
			fmt.Println("Producer stopped:", id)
			return
		case jobs <- job:
			fmt.Println("Produced:", job)
		}
	}
}

func consumer(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Consumer stopped:", id)
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			fmt.Println("Consumer", id, "processed", job)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int, 5)

	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup

	for i := 1; i <= 3; i++ {
		producerWG.Add(1)
		go producer(ctx, i, jobs, &producerWG)
	}

	for i := 1; i <= 2; i++ {
		consumerWG.Add(1)
		go consumer(ctx, i, jobs, &consumerWG)
	}

	time.AfterFunc(2*time.Second, cancel)

	producerWG.Wait()
	close(jobs)

	consumerWG.Wait()
}