package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Real example: limit concurrent external API calls to 3.
	semaphore := NewSemaphore(3)

	var wg sync.WaitGroup
	totalRequests := 10

	// It will create 10 totalRequest go routine instantly
	for requestID := 1; requestID <= totalRequests; requestID++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			semaphore.Wait()
			defer semaphore.Signal()

			fmt.Printf("Request %d started\n", id)
			time.Sleep(800 * time.Millisecond)
			fmt.Printf("Request %d completed\n", id)
		}(requestID)
	}

	wg.Wait()
	fmt.Println("All requests processed with max 3 concurrent calls")
}
