package main

import (
	"fmt"
	"time"
)

func main() {

	// -----------------------------------
	// 1. Simple switch
	// Go automatically breaks after each case
	// -----------------------------------
	i := 3
	switch i {
	case 1:
		fmt.Println("one") // no need to write break in Go
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	// -----------------------------------
	// 2. Multiple values in a case
	// -----------------------------------
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// -----------------------------------
	// 3. Type switch (detects data type)
	// -----------------------------------
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("other type")
		}
	}

	whoAmI(55)
	whoAmI("golang")
	whoAmI(true)

	// -----------------------------------
	// 4. Type switch with variable
	// -----------------------------------
	whoAmI2 := func(i interface{}) {
		switch t := i.(type) { // t holds the actual value
		case int:
			fmt.Println("integer:", t)
		case string:
			fmt.Println("string:", t)
		case bool:
			fmt.Println("boolean:", t)
		default:
			fmt.Println("other:", t)
		}
	}

	whoAmI2(99.5)
}
