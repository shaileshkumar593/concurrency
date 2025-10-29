package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	divide(a, b)

}

func divide(a int, b int) {
	fmt.Printf("%d / %d = %d", a, b, a/b)

}

func main() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}
