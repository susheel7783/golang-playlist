package main

import "fmt"

// Array = numbered sequence of elements of a fixed length

func main() {

	// -----------------------------------
	// 1. Zero-value array
	// int -> 0, string -> "", bool -> false
	// -----------------------------------
	var nums [4]int // declare an array of length 4

	fmt.Println(nums)        // [0 0 0 0]
	fmt.Println("Length:", len(nums))

	// assign value
	nums[0] = 1
	fmt.Println("First value:", nums[0])
	fmt.Println("Full array:", nums)

	// -----------------------------------
	// 2. Boolean array
	// -----------------------------------
	var vals [4]bool
	vals[2] = true
	fmt.Println("Bool array:", vals)

	// -----------------------------------
	// 3. String array
	// -----------------------------------
	var names [3]string
	names[0] = "golang"
	fmt.Println("String array:", names)

	// -----------------------------------
	// 4. Declare and initialize in one line
	// -----------------------------------
	numbers := [3]int{1, 2, 3}
	fmt.Println("One line array:", numbers)

	// -----------------------------------
	// 5. 2D array (matrix)
	// -----------------------------------
	matrix := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println("2D array:", matrix)

}
------------
ey concepts you wrote (corrected and refined)

✅ Fixed size → length cannot change

✅ Predictable memory layout

✅ Constant time access (O(1)) using index

✅ Zero values by default
