package main

import "fmt"

// -----------------------------------
// By value (copy is passed)
// This will NOT change the original value
// -----------------------------------
func changeByValue(num int) {
	num = 5
	fmt.Println("Inside changeByValue:", num)
}

// -----------------------------------
// By reference (address is passed)
// This WILL change the original value
// -----------------------------------
func changeByReference(num *int) {
	*num = 5 // dereferencing pointer
	fmt.Println("Inside changeByReference:", *num)
}

func main() {

	num := 1

	// -----------------------------
	// Pass by value
	// -----------------------------
	changeByValue(num)
	fmt.Println("After changeByValue:", num) // still 1

	// -----------------------------
	// Pass by reference (pointer)
	// -----------------------------
	changeByReference(&num)
	fmt.Println("After changeByReference:", num) // now 5

	// -----------------------------
	// Memory address
	// -----------------------------
	fmt.Println("Memory address of num:", &num)
}

--------------------------
Key pointer concepts (interview-ready)

&num → gives address of variable

*num → dereference (get/set value at that address)

By default, Go passes copy of value

Use pointers when you want to:

modify original data

avoid copying large structures

share memory

🎯 Very important understanding
func changeByReference(num *int)


Means:

num is not the value, num is the address where the value is stored.

*num = 5


Means:

go to that address and update the value.

🚀 Real backend usage

Pointers are used everywhere in Go:

struct methods

database models

JSON decoding

performance optimization

shared state
