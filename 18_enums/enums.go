package main

import "fmt"

// ENUM STYLE IN GO
// We create a new custom type based on string.
// This improves type safety and code clarity.
type OrderStatus string

// Fixed set of allowed values (enum simulation)
const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

// This function only accepts OrderStatus type
// instead of raw string to avoid invalid values.
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

func main() {

	// Only predefined constants should be used
	changeOrderStatus(Prepared)

	// This is allowed but discouraged:
	// changeOrderStatus("abc")

	// Best practice: always use enum constants.
}
------------------------------------------
Go doesn’t have native enums, so we use custom types with constants to achieve type safety, readability, and to prevent invalid values.
