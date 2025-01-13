package main

import (
	"fmt"
	"time"
)

// The main function demonstrates how goroutines work in Go.
// We trigger a goroutine using the `go` keyword.
// Goroutines allow concurrent execution, but the main program does not wait for them to finish unless explicitly managed.
func main() {
	go greeter("Hello") // Launches the `greeter` function as a goroutine.
	greeter("World")    // Executes the `greeter` function in the main thread.
}

// greeter prints the given string 5 times with a small delay in between.
// This simulates some concurrent task.
func greeter(s string) {
	for i := 0; i < 5; i++ {
		// Sleep for 3 milliseconds to simulate a time-consuming task.
		// This also helps demonstrate how goroutines allow other tasks to run concurrently.
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s) // Print the string passed to the function.
	}
}
