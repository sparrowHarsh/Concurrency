package main

import "sync"

// mutext for protecting a simple counter varibale which has been incremented by n number of threads

type Counter struct {
	mu      sync.Mutex
	counter int
}

// Increment function (Defer make sure that block is always released even if function panics)
func (c *Counter) IncrementWithDefer() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// Using manual lock
func (c *Counter) IncrementManula() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

// using TryLock -> Always preffered
func (c *Counter) Increment() bool {
	// If available it will acquire the lock
	if c.mu.TryLock() {
		defer c.mu.Unlock()
		c.counter++
		return true
	}
	return false
}
