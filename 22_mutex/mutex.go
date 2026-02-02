package main

import (
	"fmt"
	"sync"
)

// post represents a shared resource
// views -> shared variable (accessed by multiple goroutines)
// mu    -> Mutex used to protect views from race conditions
type post struct {
	views int
	mu    sync.Mutex
}

// inc increments the views count safely
// p  -> pointer receiver so changes affect original struct
// wg -> WaitGroup to synchronize goroutines
func (p *post) inc(wg *sync.WaitGroup) {

	// defer ensures these run when function exits
	// 1️⃣ Unlock the mutex
	// 2️⃣ Tell WaitGroup this goroutine is done
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	// Lock the mutex BEFORE accessing shared data
	// Only ONE goroutine can pass this line at a time
	p.mu.Lock()

	// Critical section (shared memory access)
	p.views += 1
}

func main() {

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create post with initial views = 0
	myPost := post{views: 0}

	// Start 100 goroutines
	for i := 0; i < 100; i++ {

		// Increment WaitGroup counter
		wg.Add(1)

		// Each goroutine increments views safely
		go myPost.inc(&wg)
	}

	// Block main until all goroutines call wg.Done()
	wg.Wait()

	// Safe final value (will always be 100)
	fmt.Println(myPost.views)
}
'-------------------------------------
Multiple goroutines are trying to update the same variable (views).

Without protection:

Data gets corrupted

Final value will be less than 100

This is called a race condition

🔒 What is sync.Mutex?

A Mutex is a lock.

Think of it like:

“Only one person can enter this room at a time.”

p.mu.Lock()   // enter room
p.mu.Unlock() // leave room

🧵 What happens when program runs?

1️⃣ 100 goroutines start
2️⃣ Each goroutine calls inc()
3️⃣ Mutex allows only one goroutine to update views
4️⃣ Others wait until lock is released
5️⃣ Final value becomes 100 (correct)

⚠️ What happens WITHOUT mutex?
p.views += 1


This is actually:

read value

add 1

write value

Multiple goroutines can overlap → ❌ wrong result

🧠 Why pointer receiver *post?
func (p *post) inc()


Because:

We want to modify the original post

Not a copy

Mutex must also be shared

🧠 Interview one-liners

Mutex: “Used to protect shared memory from concurrent access”

Race condition: “When multiple goroutines access shared data without synchronization”
