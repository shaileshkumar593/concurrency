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
	"sync"
	"testing"
)

type Inventory struct {
	Qty    int
	ItemId string
}

var mu sync.Mutex

func NewInventory(initialStock map[string]int) *Inventory {
	var s *Inventory
	mu.Lock()
	for key, val := range initialStock {
		s.Qty = val
		s.ItemId = key
	}
	mu.Unlock()

	return s

}

func TestNewInventory(t *testing.T) {
	aa := map[string]int{
		"toy": 50,
	}
	p := NewInventory(map[string]int{
		"toy": 50})

	if p == nil {
		t.Error("Inventory not allocated properly")
	}

	if p.Qty == aa["toy"] {
		t.Log("inventory allocated successfully")
	}

	if p.Qty < aa["toy"] || p.Qty > aa["toy"] {
		t.Log("wrong Inventory allocation ")
	}

}

func (i *Inventory) Reserve(itemID string, qty int) bool {
	mu.Lock()
	var status bool
	if i.Qty > qty {
		status = true
		i.Qty = i.Qty - qty
	} else {
		status = false
	}
	mu.Unlock()

	return status
}

func TestReserve(t *testing.T) {

	p := NewInventory(map[string]int{})

	status := p.Reserve("toy", 10)

	if status == true {
		t.Log("successfully reserved ")
	} else {
		t.Log("Wrong allocation ")
	}

}

// Release returns `qty` units of `itemID` back to inventory.

func (i *Inventory) Release(itemID string, qty int) int {

	mu.Lock()
	i.Qty = i.Qty + qty
	mu.Unlock()
	return qty
}

func TestRelease(t *testing.T) {

	p := NewInventory(map[string]int{})

	qty := p.Release("toy", 3)

	if qty == 3 {
		t.Logf("Released %d resources ", qty)
	} else {
		t.Log("Wrong release ")
	}

}

func main() {

	m := map[string]int{
		"toy": 10,
	}

	inventory := NewInventory(m)
	for i := 0; i < 3; i++ {

		go inventory.Reserve("toy", 3)
		go inventory.Release("toy", 3)
	}
}
