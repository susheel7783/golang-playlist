package main

import (
	"fmt"
	"maps" // Go 1.21+
)

// maps -> hash table (like dict in Python, object in JS)

func main() {

	// -----------------------------------
	// 1. Creating a map using make()
	// key: string, value: string
	// -----------------------------------
	m := make(map[string]string)

	// setting elements
	m["name"] = "golang"
	m["area"] = "backend"

	// getting elements
	fmt.Println("name:", m["name"])
	fmt.Println("area:", m["area"])

	// if key does NOT exist, zero value is returned
	fmt.Println("unknown:", m["unknown"]) // ""

	// -----------------------------------
	// 2. Map with int values
	// -----------------------------------
	m2 := make(map[string]int)
	m2["age"] = 30
	m2["price"] = 50

	fmt.Println("\nm2:", m2)
	fmt.Println("len:", len(m2))

	// -----------------------------------
	// 3. delete() and clear()
	// -----------------------------------
	delete(m2, "price") // removes one key
	fmt.Println("after delete:", m2)

	clear(m2) // removes all keys
	fmt.Println("after clear:", m2)

	// -----------------------------------
	// 4. Creating map without make()
	// -----------------------------------
	m3 := map[string]int{"price": 40, "phones": 3}
	fmt.Println("\nm3:", m3)

	// -----------------------------------
	// 5. Check if key exists
	// -----------------------------------
	v, ok := m3["phones"] // v = value, ok = true/false
	if ok {
		fmt.Println("phones exists, value =", v)
	} else {
		fmt.Println("phones key not found")
	}

	// -----------------------------------
	// 6. Compare two maps (Go 1.21+)
	// -----------------------------------
	m4 := map[string]int{"price": 40, "phones": 3}
	m5 := map[string]int{"price": 40, "phones": 8}

	fmt.Println("\nmaps equal:", maps.Equal(m4, m5)) // false
}
-------------------------
Maps are reference types

Accessing missing key → returns zero value

Use value, ok := m[key] to check existence

delete(m, key) → removes one entry

clear(m) → empties entire map

maps.Equal() → compares maps safely
