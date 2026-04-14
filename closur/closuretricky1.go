package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	// Give goroutines time to execute
	time.Sleep(time.Second)
}