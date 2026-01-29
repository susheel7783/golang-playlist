package main

import "fmt"

/*
INTERFACE:
An interface defines WHAT a type can do, not HOW it does it.

Any struct that implements all these methods
automatically satisfies this interface.
*/
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

/*
This struct depends on the interface, not a real payment gateway. 

WHY WE USE INTERFACE HERE:
- To avoid tight coupling
- To easily switch payment providers
- To follow Open/Closed Principle
- To make testing easy (fake gateway)
*/
type payment struct {
	gateway paymenter
}

// This method works with ANY payment gateway
// as long as it follows the interface
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

func (p payment) makeRefund(amount float32, account string) {
	p.gateway.refund(amount, account)
}

/*
Razorpay is a real implementation.
Because it has both methods, it automatically
becomes a paymenter.
*/
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay payment:", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("Razorpay refund:", amount, "to", account)
}

/*
Fake implementation for testing.
No real API calls.
*/
type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("Fake payment:", amount)
}

func (f fakepayment) refund(amount float32, account string) {
	fmt.Println("Fake refund:", amount)
}

func main() {

	// You can plug in ANY gateway
	fakeGw := fakepayment{}

	newPayment := payment{
		gateway: fakeGw, // depends on interface, not real object
	}

	newPayment.makePayment(100)
	newPayment.makeRefund(50, "test@upi")
}
---------------------------------------------
 ## interface 
An interface in Go defines behavior. We use it to achieve abstraction, loose coupling, easy testing, and to follow the Open/Closed principle
