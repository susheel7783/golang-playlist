package main

import "fmt"

// const age = 30   // package-level constant (outside main)

func main() {

	// -----------------------------------
	// Single constant declaration
	// -----------------------------------
	const name = "golang"
	const age = 30

	fmt.Println(name, age)

	// -----------------------------------
	// Grouped constant declaration
	// Use this when you have a list of constants
	// -----------------------------------
	const (
		port = 5000
		host = "localhost"
		debug = true
	)

	fmt.Println(port, host, debug)
}
-----------
✅ Valid
const age = 30
const pi = 3.14

const (
	app = "myApp"
	version = 1
)

❌ Invalid
const age := 30     // ❌ shorthand not allowed with const

const age          // ❌ value is mandatory

🎯 Key points

const values are fixed (cannot be changed)

const must be assigned at declaration

:= works only with var, not with const

Grouped const ( ... ) is used for clean code and configs
