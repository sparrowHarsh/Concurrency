package main

// This is the implementation of semaphore

// we are using buffered channel for the semaphore
// channel capacity = permit count
// sending = wait(), receiving = signal()

type Semaphore chan struct{}

// create new semaphore
func NewSemaphore(n int) Semaphore {
	semaphore := make(chan struct{}, n)
	return semaphore
}

// Wait() function
func (s Semaphore) Wait() {
	// send to the channel, block if channel buffer is full
	s <- struct{}{}
}

// Signal makes sure that there won't be any recieve call before send by using mutex in it's internal implementation

// signal() , receive from channel
func (s Semaphore) Signal() {
	// receive channel, return a permit to semaphore
	<-s
}

// tryWait()
// If it is free then acquire it or else not block
func (s Semaphore) TryWait() bool {
	select {
	case s <- struct{}{}:
		return true
	default:
		return false
	}
}
