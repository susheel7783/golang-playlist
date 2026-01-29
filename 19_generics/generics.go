package main

import "fmt"

// GENERIC FUNCTION
// T and V are type parameters (placeholders).
// T comparable → T must support == and !=
// V string     → V must be of type string
func printSlice[T comparable, V string](items []T, name V) {

	// Loop over slice of any type
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// Example of a generic data structure (STACK - LIFO)
/*
type stack[T any] struct {
	elements []T
}
*/

func main() {

	// Boolean slice
	vals := []bool{true, false, true}

	// T → bool, V → string
	printSlice(vals, "john")

	// Other valid examples:
	// nums := []int{1, 2, 3}
	// printSlice(nums, "numbers")

	// names := []string{"golang", "typescript"}
	// printSlice(names, "languages")
}
------------------------
Generics allow us to write reusable, type-safe code by using type parameters instead of concrete types, avoiding duplication and runtime type casting
