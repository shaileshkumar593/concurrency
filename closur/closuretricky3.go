package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		i := i // create new variable
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second)
}