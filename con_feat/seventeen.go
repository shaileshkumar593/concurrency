17. RWMutex Pattern
Use case
Many readers, few writers.

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *SafeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	return val, ok
}

func main() {
	sm := NewSafeMap()

	sm.Set("a", 10)

	val, ok := sm.Get("a")
	fmt.Println(val, ok)
}
Pros
Better for read-heavy workloads

Multiple readers allowed

Cons
Writers block readers

Not always faster than Mutex

More complex