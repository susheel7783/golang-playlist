package main // marks the application's entry point; only this package can be compiled and run, and execution starts from func main()

import "fmt"

func main() {

	// ------------------------------------------------
	// 1. fmt.Print()
	// Prints text WITHOUT a new line at the end
	// ------------------------------------------------
	fmt.Print("Hello")
	fmt.Print(" World") // Output: Hello World

	fmt.Print("\n") // manually adding new line

	// ------------------------------------------------
	// 2. fmt.Println()
	// Prints text WITH a new line at the end
	// Automatically adds spaces between values
	// ------------------------------------------------
	fmt.Println("Hello")
	fmt.Println("Hello", "World", 2026)

	// ------------------------------------------------
	// 3. fmt.Printf()
	// Prints formatted output (does NOT add new line automatically)
	// %s = string, %d = integer, %f = float, %t = boolean
	// ------------------------------------------------
	name := "Susheel"
	age := 22
	marks := 91.5
	pass := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Marks: %.2f\n", marks) // .2 means 2 digits after decimal
	fmt.Printf("Pass: %t\n", pass)

	// ------------------------------------------------
	// 4. fmt.Sprint()
	// Does NOT print
	// Returns a string which we can store in a variable
	// ------------------------------------------------
	message := fmt.Sprint("Go", " is", " powerful")
	fmt.Println(message)

	// ------------------------------------------------
	// 5. fmt.Sprintf()
	// Like Printf but RETURNS formatted string instead of printing
	// ------------------------------------------------
	info := fmt.Sprintf("User %s is %d years old", name, age)
	fmt.Println(info)

	// ------------------------------------------------
	// 6. fmt.Sprintln()
	// Like Println but RETURNS a string with new line
	// ------------------------------------------------
	line := fmt.Sprintln("Learning", "Golang")
	fmt.Print(line)

	// ------------------------------------------------
	// 7. Printing any type using %v
	// ------------------------------------------------
	fmt.Printf("Any value print: %v\n", 100)
	fmt.Printf("Any value print: %v\n", 99.9)
	fmt.Printf("Any value print: %v\n", []int{1, 2, 3})

}

----------------------
Print → no new line

Println → new line

Printf → formatted printing

Sprint / Sprintf / Sprintln → don’t print, they return string
