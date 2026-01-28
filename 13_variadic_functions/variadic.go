package main

import "fmt"

// sum is a variadic function
// nums ...int means: it can take 0, 1, or many int values
func sum(nums ...int) int {

	// nums behaves like a slice []int inside the function
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {

	// -----------------------------------
	// 1. fmt.Println is itself a variadic function
	// It can take any number of values of any type
	// -----------------------------------
	fmt.Println(1, 2, 3, 4, 5, 66, "Hi")

	// -----------------------------------
	// 2. Calling variadic function with normal values
	// -----------------------------------
	result1 := sum(1, 2, 3, 6, 7)
	fmt.Println("result1:", result1)

	// -----------------------------------
	// 3. Passing a slice to a variadic function
	// Use ... to "unpack" the slice
	// -----------------------------------
	nums := []int{3, 4, 5, 6}
	result2 := sum(nums...)
	fmt.Println("result2:", result2)
}
// ----------------------
✅ Variadic function syntax
func sum(nums ...int)


Inside the function:

nums is of type []int

✅ Normal call
sum(1, 2, 3)

✅ Slice call (very important)
arr := []int{1, 2, 3}
sum(arr...)

✅ Built-in example
fmt.Println(1, "hi", true, 9.8)


Println is a variadic function.
