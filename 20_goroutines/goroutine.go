package main

import (
	"fmt"
	"sync"
)

// task function
// id  -> identifies the task
// w   -> pointer to WaitGroup so all goroutines share the same counter
func task(id int, w *sync.WaitGroup) {

	// Done() tells the WaitGroup that this goroutine has finished
	// defer ensures Done() is called even if function exits early
	defer w.Done()

	fmt.Println("doing task", id)
}

func main() {

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Launch multiple goroutines
	for i := 0; i <= 10; i++ {

		// Increment WaitGroup counter
		// Must be called BEFORE starting goroutine
		wg.Add(1)

		// Start goroutine
		go task(i, &wg)
	}

	// Wait blocks main goroutine
	// until WaitGroup counter becomes zero
	wg.Wait()

	// Program exits only after all goroutines finish
}
-----------------------------------------
WaitGroup is used to synchronize goroutines. It allows the main goroutine to wait until a set of spawned goroutines completes execution.
