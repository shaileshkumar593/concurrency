package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 10
	chOdd := make(chan bool)
	chEven := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(n, chOdd, chEven, &wg)
	go even(n, chOdd, chEven, &wg)
	wg.Wait()
}

func odd(n int, chOdd chan bool, chEven chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i = i + 2 {
		fmt.Println(i)
		chEven <- true
		<-chOdd
	}
}

func even(n int, chOdd chan bool, chEven chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= n; i = i + 2 {
		<-chEven
		fmt.Println(i)
		chOdd <- true
	}
}
