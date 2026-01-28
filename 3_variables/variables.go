package main

import "fmt"

func main() {

	// -----------------------------------
	// 1. Explicit type declaration
	// You clearly tell Go the data type
	// -----------------------------------
	var name string = "golang"
	var isAdult bool = true
	var age int = 30
	var price float32 = 50.5

	fmt.Println(name, isAdult, age, price)

	// -----------------------------------
	// 2. Type inference
	// Go automatically understands the type
	// -----------------------------------
	var course = "golang"   // string
	var passed = true      // bool
	var year = 2026        // int
	var cost = 99.99       // float64 (default float type)

	fmt.Println(course, passed, year, cost)

	// -----------------------------------
	// 3. Shorthand syntax (:=)
	// Only works INSIDE functions
	// Most commonly used in real projects
	// -----------------------------------
	language := "Go"
	active := true
	count := 100
	rating := 4.8

	fmt.Println(language, active, count, rating)

	// -----------------------------------
	// 4. Declare first, assign later
	// Here you MUST write the type
	// -----------------------------------
	var company string
	company = "Tilda"

	var salary int
	salary = 50000

	fmt.Println(company, salary)
}
-------------------
✅ You must write type when:
var name string

✅ You can skip type when value is given:
var name = "golang"

✅ Short form (most used):
name := "golang"

❌ This is invalid in Go:
var name        // type missing

🎯 Interview-level one-liners

var name string = "go" → explicit declaration

var name = "go" → type inference

name := "go" → shorthand, only inside functions
