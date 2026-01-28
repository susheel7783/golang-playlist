package main

import (
	"fmt"
	"slices" // Go 1.21+
)

// slice -> dynamic array (can grow and shrink)
// this is the most used data structure in Go

func main() {

	// -----------------------------------
	// 1. Uninitialized slice (nil slice)
	// -----------------------------------
	var nums []int

	fmt.Println(nums == nil) // true (nil slice)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))

	// -----------------------------------
	// 2. make() - initialize slice (not nil)
	// make(type, length, capacity)
	// -----------------------------------
	nums = make([]int, 0, 5)

	fmt.Println("\nAfter make:")
	fmt.Println(nums == nil) // false
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))

	// -----------------------------------
	// 3. append() - add elements
	// -----------------------------------
	nums = append(nums, 1)
	nums = append(nums, 2, 3)

	fmt.Println("\nAfter append:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))

	// -----------------------------------
	// 4. make with length and capacity
	// -----------------------------------
	nums2 := make([]int, 2, 5)
	nums2[0] = 3
	nums2[1] = 5

	fmt.Println("\nnums2:", nums2)
	fmt.Println("len:", len(nums2))
	fmt.Println("cap:", cap(nums2))

	// -----------------------------------
	// 5. copy() - copy one slice to another
	// -----------------------------------
	nums3 := make([]int, len(nums))
	copy(nums3, nums)

	fmt.Println("\nAfter copy:")
	fmt.Println("nums :", nums)
	fmt.Println("nums3:", nums3)

	// -----------------------------------
	// 6. Slice operator
	// -----------------------------------
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("\nSlicing examples:")
	fmt.Println(arr[0:1]) // [1]
	fmt.Println(arr[:2])  // [1 2]
	fmt.Println(arr[1:])  // [2 3 4 5]

	// -----------------------------------
	// 7. Compare slices (Go 1.21+)
	// -----------------------------------
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}

	fmt.Println("\nSlices equal:", slices.Equal(a, b)) // false

	// -----------------------------------
	// 8. 2D slice
	// -----------------------------------
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("\n2D slice:", matrix)
}
---------------
Slice is a dynamic view over an array

nil slice ≠ empty slice

len() → number of elements

cap() → how many elements it can grow without reallocating

append() → automatically grows capacity

copy() → copies values, not reference

slices.Equal() → safe slice comparison
