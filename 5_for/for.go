package main

import "fmt"

// In Go, `for` is the ONLY looping construct.
// There is no separate while or do-while loop.

func main() {

	// -----------------------------------
	// 1. While-style loop using `for`
	// -----------------------------------
	i := 1
	for i <= 3 {
		fmt.Println("while-style:", i)
		i = i + 1
	}

	// -----------------------------------
	// 2. Infinite loop
	// -----------------------------------
	// for {
	// 	fmt.Println("running forever")
	// }

	// -----------------------------------
	// 3. Classic for loop
	// -----------------------------------
	for j := 0; j <= 3; j++ {

		// break example
		// if j == 3 {
		// 	break
		// }

		// continue example (skips 2)
		if j == 2 {
			continue
		}

		fmt.Println("classic:", j)
	}

	// -----------------------------------
	// 4. Range-based loop (Go 1.22+ feature)
	// -----------------------------------
	// This prints numbers from 0 to 10
	for k := range 11 {
		fmt.Println("range:", k)
	}

}
--------------
🧠 Important points (interview ready)

Go has only one loop keyword → for

for condition {} → works like while

for {} → infinite loop

for init; condition; inc {} → classic loop

for i := range 11 {} → Go 1.22+, loops from 0 to n-1
