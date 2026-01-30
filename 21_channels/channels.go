package main

import (
	"fmt"
	"time"
)

/*
==================================================
BASIC IDEA (VERY IMPORTANT)
==================================================

1️⃣ Goroutine
   - Lightweight thread managed by Go
   - Created using the `go` keyword

2️⃣ Channel
   - A pipe used by goroutines to communicate
   - One goroutine sends data
   - Another goroutine receives data

3️⃣ Channels are BLOCKING by default
   - Send waits for receive
   - Receive waits for send

==================================================
*/

// ------------------------------------------------
// emailSender FUNCTION
// ------------------------------------------------

// emailChan <-chan string
//   - receive-only channel
//   - this function can ONLY read from this channel
//
// done chan<- bool
//   - send-only channel
//   - this function can ONLY send to this channel
func emailSender(emailChan <-chan string, done chan<- bool) {

	// defer means:
	// "Run this line when the function finishes"
	// Here: we are signaling that work is done
	defer func() { done <- true }()

	// range keeps receiving values from the channel
	// It STOPS only when the channel is closed
	for email := range emailChan {

		// Processing email
		fmt.Println("sending email to", email)

		// Simulate delay (1 second)
		time.Sleep(time.Second)
	}
}

func main() {

	// ==================================================
	// MULTIPLE CHANNELS + SELECT
	// ==================================================

	// Create an unbuffered channel of type int
	// Unbuffered means:
	// sender waits until receiver is ready
	chan1 := make(chan int)

	// Create an unbuffered channel of type string
	chan2 := make(chan string)

	// Start a goroutine that sends value to chan1
	go func() {
		// This will BLOCK until someone receives
		chan1 <- 10
	}()

	// Start another goroutine that sends value to chan2
	go func() {
		chan2 <- "pong"
	}()

	// We know 2 messages will come (one from each channel)
	for i := 0; i < 2; i++ {

		// select waits for ANY channel to be ready
		// Whichever channel sends first will execute
		select {

		// Receive value from chan1
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		// Receive value from chan2
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// ==================================================
	// EMAIL WORKER EXAMPLE (BUFFERED CHANNEL)
	// ==================================================

	/*
	// Buffered channel with capacity 100
	// Can hold 100 emails without blocking sender
	emailChan := make(chan string, 100)

	// Channel used only for synchronization
	done := make(chan bool)

	// Start email sender goroutine
	go emailSender(emailChan, done)

	// Send emails into the channel
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending emails")

	// VERY IMPORTANT:
	// close() tells receiver:
	// "No more values will be sent"
	close(emailChan)

	// Block main until emailSender finishes
	<-done
	*/

	// ==================================================
	// SIGNAL CHANNEL (GOROUTINE SYNCHRONIZATION)
	// ==================================================

	/*
	done := make(chan bool)

	go func() {
		fmt.Println("processing...")
		done <- true // signal completion
	}()

	// main waits here until signal is received
	<-done
	*/

	// ==================================================
	// CHANNEL USED FOR RETURNING VALUES
	// ==================================================

	/*
	result := make(chan int)

	go func() {
		result <- 4 + 5 // send result
	}()

	// Receive result (blocking)
	res := <-result
	fmt.Println(res)
	*/

	// ==================================================
	// BLOCKING BEHAVIOR (IMPORTANT WARNING)
	// ==================================================

	/*
	messageChan := make(chan string)

	// This line will BLOCK FOREVER
	// because no goroutine is receiving
	messageChan <- "ping"

	msg := <-messageChan
	fmt.Println(msg)
	*/
}
--------------------------------
Goroutine → runs code concurrently

Channel → pipe for communication

Send (<-) → blocks until receive

Receive (<-) → blocks until send

select → listen to multiple channels

close(channel) → no more data

Buffered channel → works like a queue
