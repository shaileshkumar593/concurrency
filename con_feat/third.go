// 3. Channel Communication Pattern
/*Use case
One goroutine sends data, another receives data.*/


package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}

/*
Pros
	Safe communication

	Avoids shared memory

Cons
	Can deadlock if no receiver

	Requires careful close handling
*/
