package main

import (
	"fmt"
	"time"
)

/*
--------------------------------------------------
CHANNEL CONCEPTS USED IN THIS FILE
--------------------------------------------------
1. Sending data to a channel
2. Receiving data from a channel
3. Direction-only channels
4. select statement
5. Buffered channels
6. close(channel)
7. Goroutine synchronization using channels
--------------------------------------------------
*/

// emailSender demonstrates:
// - receive-only channel for emails
// - send-only channel for signaling completion
func emailSender(emailChan <-chan string, done chan<- bool) {

	// Ensure done signal is sent when function exits
	defer func() { done <- true }()

	// range keeps receiving values until channel is closed
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	// ---------------------------------------------
	// MULTIPLE CHANNELS + SELECT
	// ---------------------------------------------

	// Create unbuffered channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	// Send value to chan1 using goroutine
	go func() {
		chan1 <- 10
	}()

	// Send value to chan2 using goroutine
	go func() {
		chan2 <- "pong"
	}()

	// select waits for ANY channel to be ready
	// It executes whichever case is available first
	for i := 0; i < 2; i++ {
		select {

		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// ---------------------------------------------
	// EMAIL WORKER EXAMPLE (BUFFERED CHANNEL)
	// ---------------------------------------------

	/*
	emailChan := make(chan string, 100) // buffered channel
	done := make(chan bool)             // synchronization channel

	// Start email sender goroutine
	go emailSender(emailChan, done)

	// Send emails
	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending emails")

	// VERY IMPORTANT:
	// close tells receiver that no more data will come
	close(emailChan)

	// Block main until emailSender finishes
	<-done
	*/

	// ---------------------------------------------
	// SIGNAL CHANNEL (GOROUTINE SYNCHRONIZATION)
	// ---------------------------------------------

	/*
	done := make(chan bool)

	go func() {
		fmt.Println("processing...")
		done <- true
	}()

	<-done // blocks main until goroutine finishes
	*/

	// ---------------------------------------------
	// CHANNEL USED FOR RETURNING VALUES
	// ---------------------------------------------

	/*
	result := make(chan int)

	go func() {
		result <- 4 + 5
	}()

	res := <-result // blocks until value received
	fmt.Println(res)
	*/

	// ---------------------------------------------
	// BLOCKING BEHAVIOR DEMO
	// ---------------------------------------------

	/*
	messageChan := make(chan string)

	// This will block forever if no receiver exists
	messageChan <- "ping"

	msg := <-messageChan
	fmt.Println(msg)
	*/
}
--------------------------------------------
Channel

Used to communicate between goroutines safely

🔹 Goroutine

Lightweight thread managed by Go runtime

🔹 select

Waits on multiple channel operations and executes the first available one

🔹 close(channel)

Signals that no more values will be sent

🔹 Buffered channel

Allows sending without immediate receiver (up to capacity)


🧠 Interview one-liner

“Channels allow goroutines to communicate and synchronize by passing data, following the principle: don’t communicate by sharing memory, share memory by communicating.”
