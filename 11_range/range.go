package main

import "fmt"

// Iterating over data structures using for and range

func main() {

	// -----------------------------------
	// 1. Iterating over a slice (classic loop)
	// -----------------------------------
	nums := []int{6, 7, 8}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i]) // 6 7 8
	}

	// -----------------------------------
	// 2. Iterating over a slice using range
	// -----------------------------------
	for _, num := range nums {
		fmt.Println(num) // values only
	}

	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	// -----------------------------------
	// 3. Sum of elements
	// -----------------------------------
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println("sum:", sum) // 21

	// -----------------------------------
	// 4. Iterating over a map
	// -----------------------------------
	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	for k := range m { // only keys
		fmt.Println("key only:", k)
	}

	// -----------------------------------
	// 5. Iterating over a string (runes / Unicode)
	// -----------------------------------
	// range over string gives:
	// i = starting byte index
	// c = rune (Unicode code point)

	for i, c := range "golang" {
		fmt.Println(i, c) // Unicode value
	}

	for i, c := range "golang" {
		fmt.Println(i, string(c)) // actual character
	}
}
// --------------------------
🔹 Slice / Array
for i, v := range nums


i → index

v → value

🔹 Map
for k, v := range m


order is not guaranteed

🔹 String
for i, c := range "golang"


i → byte index

c → rune (Unicode code point)
