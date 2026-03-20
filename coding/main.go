package main

/*
Go Coding Exercise: Inventory Reservation Service

Business Context

In many backend systems (e-commerce, ticketing, logistics), inventory must be reserved during a

workflow (e.g. checkout) before final confirmation (e.g. payment).

The system must ensure:

• Inventory is never over-reserved

• Concurrent requests do not corrupt state

• Failures do not result in negative stock

Problem Statement

Implement a concurrency-safe, in-memory inventory reservation service.

type Inventory struct {

// your fields here

}

func NewInventory(initialStock map[string]int) *Inventory

// Reserve attempts to reserve `qty` units of `itemID`.

// Returns true if reservation succeeds.

// Returns false if insufficient stock.

func (i *Inventory) Reserve(itemID string, qty int) bool

// Release returns `qty` units of `itemID` back to inventory.

func (i *Inventory) Release(itemID string, qty int)

Functional Requirements

• Inventory starts with an initial stock per item

• Reserve must:

o Be atomic

o Never allow stock to go below zero

• Release must:

o Return stock safely

o Never panic

• The implementation must be safe under concurrent access

Constraints (Read Carefully)

• In-memory only (no DB, no Redis)

• Standard library only

• Focus on correctness and clarity

• Assume this code will be used in a multi-goroutine backend service

You do not need to implement everything below — they are guidance, not

requirements.

• Think about race conditions when two goroutines reserve the same item

• A single process can still receive concurrent requests

• Correctness is more important than performance

• A simple locking strategy is perfectly acceptable

• You may assume:

o Missing items have zero stock

o Invalid quantities (≤ 0) can be safely ignored or rejected
*/

import (
	"fmt"
	"sync"
)

type Inventory struct {
	mu    sync.Mutex
	stock map[string]int
}

// Constructor
func NewInventory(initialStock map[string]int) *Inventory {
	// copy map to avoid external mutation
	stockCopy := make(map[string]int)
	for k, v := range initialStock {
		stockCopy[k] = v
	}

	return &Inventory{
		stock: stockCopy,
	}
}

// Reserve attempts to reserve qty units
func (i *Inventory) Reserve(itemID string, qty int) bool {
	if qty <= 0 {
		return false
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	current := i.stock[itemID]

	if current < qty {
		return false
	}

	i.stock[itemID] = current - qty
	return true
}

// Release adds qty back to inventory
func (i *Inventory) Release(itemID string, qty int) {
	if qty <= 0 {
		return
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	i.stock[itemID] += qty
}

func main() {
	inv := NewInventory(map[string]int{
		"itemA": 10,
	})

	var wg sync.WaitGroup

	// 20 concurrent reservations of 1 unit each
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if inv.Reserve("itemA", 1) {
				fmt.Println("Reserved by goroutine", id)
			} else {
				fmt.Println("Failed by goroutine", id)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Final stock:", inv.stock["itemA"])
}
