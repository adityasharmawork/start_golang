package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	
	defer wg.Done()

	fmt.Printf("Worker %d started!\n", i)
	// Some task is happening here
	fmt.Printf("Worker %d completed!\n", i)

}

func main() {
	fmt.Println("Explore Goroutines")

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()
	
	fmt.Println("Workers task completed")

}
