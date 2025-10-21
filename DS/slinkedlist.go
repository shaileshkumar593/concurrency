package main

import (
	"fmt"
	"os"
)

type Node struct {
	value int
	next  *Node
}

var head *Node = nil
var last *Node = nil
var NodeCnt int

func InserNodeEnd(x int) {
	fmt.Println("Added node")
	//var cur *Node = nil
	var temp Node
	//var temp *Node = nil
	//temp =  New(Node)
	fmt.Println("----")
	temp.value = x
	temp.next = nil
	//fmt.Println("head   ", head)
	//fmt.Printf("%p\n", head)
	if head == nil {
		NodeCnt = NodeCnt + 1
		head = &temp
		last = head
	} else {
		NodeCnt = NodeCnt + 1
		last.next = &temp
		last = last.next
	}

}

func InsertFront(x int) {
	fmt.Println(" \n Added node at Frond ")
	//var cur *Node = nil
	//var temp Node
	var temp *Node = nil
	temp = new(Node)
	fmt.Println("----")
	temp.value = x
	temp.next = nil
	fmt.Println("head   ", head)
	fmt.Printf("%p\n", head)
	if head == nil {
		NodeCnt = NodeCnt + 1
		head = temp
		last = head
	} else {
		NodeCnt = NodeCnt + 1
		temp.next = head
		head = temp
	}

}

func InsertPos(x int, pos int) {
	var cur *Node = head
	var cnt int

	if pos == 0 {
		InsertFront(x)
	} else if pos == NodeCnt {
		InserNodeEnd(x)
	} else {
		var temp *Node = nil
		temp = new(Node)
		temp.value = x
		temp.next = nil
		for pos-1 != cnt {
			cnt = cnt + 1
			cur = cur.next
		}
		cur.next = temp
	}

}

func display() {
	if head == nil {
		fmt.Println("Linked List is empty")
	} else {
		fmt.Println("\n Elements of linked list are:")
		var cur *Node = head
		for cur != nil {
			fmt.Printf("%d ------->", cur.value)
			cur = cur.next
		}
	}
}

func main() {
	var choice int

	for {
		fmt.Println("Choice 1. InsertNode 2. Display 3. InsertFront 4.InserPos 5. exit")
		fmt.Println("\n Enter choice")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var x int
			fmt.Println("Enter value")
			fmt.Scan(&x)
			InserNodeEnd(x)
		case 2:
			display()

		case 3:
			var x int
			fmt.Println("Enter value")
			fmt.Scan(&x)
			InsertFront(x)

		case 4:
			var x, pos int
			fmt.Println("Enter position and value")
			fmt.Scan(&x, &pos)
			InsertPos(x, pos)

		case 5:
			os.Exit(0)

		}
	}
}
