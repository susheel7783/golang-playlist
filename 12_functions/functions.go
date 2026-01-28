package main

import "fmt"

// -----------------------------------
// 1. Normal function
// -----------------------------------
func add(a int, b int) int {
	return a + b
}

// if both parameters are same type, we can write like this:
// func add(a, b int) int { ... }

// -----------------------------------
// 2. Multiple return values
// -----------------------------------
func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

// -----------------------------------
// 3. Function as parameter (callback)
// -----------------------------------
func processIt(fn func(a int) int) {
	result := fn(1)
	fmt.Println("callback result:", result)
}

// -----------------------------------
// 4. Function returning a function
// -----------------------------------
func processIt2() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

func main() {

	// -----------------------------------
	// Anonymous function assigned to variable
	// -----------------------------------
	fn := func(a int) int {
		return a + 10
	}

	processIt(fn) // passing function as argument

	// -----------------------------------
	// Function returning function
	// -----------------------------------
	newFn := processIt2()
	fmt.Println("returned function result:", newFn(6))

	// -----------------------------------
	// Multiple return values
	// -----------------------------------
	lang1, lang2, ok := getLanguages()
	fmt.Println(lang1, lang2, ok)

	// ignore one return value using _
	l1, l2, _ := getLanguages()
	fmt.Println(l1, l2)

	// -----------------------------------
	// Normal function call
	// -----------------------------------
	result := add(3, 5)
	fmt.Println("add result:", result)
}
// ----------------
✅ Normal function
func add(a, b int) int

✅ Multiple returns
func test() (int, string, bool)

✅ Function as argument
func process(fn func(int) int)

✅ Function returning function
func getFn() func(int) int

✅ Anonymous function
fn := func(a int) int { return a }
