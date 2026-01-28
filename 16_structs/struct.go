package main

import (
	"fmt"
	"time"
)

// -----------------------------------
// Customer struct
// -----------------------------------
type customer struct {
	name  string
	phone string
}

// -----------------------------------
// Order struct (with embedding)
// customer is EMBEDDED here
// so order automatically has name and phone via customer
// -----------------------------------
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer  // embedded struct (composition)
}

// -----------------------------------
// There is no constructor in Go,
// but we usually create a "constructor-like" function
// -----------------------------------
func newOrder(id string, amount float32, status string, cust customer) *order {

	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
		customer:  cust,
	}

	return &myOrder
}

// -----------------------------------
// Method with pointer receiver
// Can modify original struct
// -----------------------------------
func (o *order) changeStatus(status string) {
	o.status = status
}

// -----------------------------------
// Method with value receiver
// Only reads data (does not modify original)
// -----------------------------------
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// -----------------------------------
	// Create customer
	// -----------------------------------
	cust := customer{
		name:  "john",
		phone: "1234567890",
	}

	// -----------------------------------
	// Create order using constructor-style function
	// -----------------------------------
	myOrder := newOrder("1", 30.5, "received", cust)

	fmt.Println("Initial order:", myOrder)

	// -----------------------------------
	// Access embedded struct fields
	// -----------------------------------
	fmt.Println("Customer name:", myOrder.name)       // direct access
	fmt.Println("Customer phone:", myOrder.phone)     // direct access
	fmt.Println("Customer name:", myOrder.customer.name) // also valid

	// -----------------------------------
	// Update embedded struct field
	// -----------------------------------
	myOrder.customer.name = "robin"

	// -----------------------------------
	// Call methods
	// -----------------------------------
	myOrder.changeStatus("confirmed")
	fmt.Println("Updated order:", myOrder)
	fmt.Println("Order amount:", myOrder.getAmount())

	// -----------------------------------
	// Inline struct (used once)
	// -----------------------------------
	language := struct {
		name   string
		isGood bool
	}{
		"golang",
		true,
	}

	fmt.Println("Inline struct:", language)
}
-----------------------------------



✅ Struct

Custom data type (like class without inheritance)

type order struct { ... }

✅ Embedding (Composition)
type order struct {
	customer
}


Now you can access:

myOrder.name
myOrder.phone


This is how Go achieves composition instead of inheritance.

✅ Constructor-style function
func newOrder(...) *order


Used to:

initialize defaults

validate data

return pointer

✅ Methods
func (o *order) changeStatus(status string)


*order → can modify original

order → read-only copy

✅ Zero values

If not set:

int → 0

float → 0

string → ""

bool → false

time.Time → 0001-01-01...
