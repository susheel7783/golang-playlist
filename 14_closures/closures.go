package main

import "fmt"

// counter does NOT take any parameters
// it RETURNS a function which returns an int
func counter() func() int {

	// this variable is captured by the inner function
	var count int = 0

	// anonymous function (closure)
	// it remembers and modifies `count`
	return func() int {
		count += 1
		return count
	}
}

func main() {

	// counter() returns a function
	// increment now holds that returned function
	increment := counter()

	fmt.Println(increment()) // output: 1
	fmt.Println(increment()) // output: 2
	fmt.Println(increment()) // output: 3

	// each call to counter() creates a NEW closure
	newCounter := counter()
	fmt.Println(newCounter()) // output: 1 (separate memory)
}
------------
What is happening here? (Very important)

counter() runs once and creates count = 0

It returns an anonymous function

That anonymous function remembers count

Every time you call increment(), it updates the same count

This remembering behavior is called a closure.

🎯 Key interview points

A closure is a function that captures variables from its outer scope

The captured variables stay alive even after the outer function finishes

Each call to counter() creates a new independent counter
