package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer

// Mutex in Go - Mutex is used to ensure mutual exclusion, protecting shared resources from concurrent access.
// It provides two methods: Lock() to acquire the lock and Unlock() to release it.
// The lock must be unlocked by the same goroutine that acquired it.
// This is required because the two or more goroutine can not write the same memory space at the same time.
var mut sync.Mutex

// The main function demonstrates how goroutines work in Go.
// We trigger a goroutine using the `go` keyword.
// Goroutines allow concurrent execution, but the main program does not wait for them to finish unless explicitly managed.
func main() {
	//go greeter("Hello") // Launches the `greeter` function as a goroutine.
	//greeter("World")    // Executes the `greeter` function in the main thread.

	// List of websites so we have to do that we have to go to these websites and get the status code out there.
	var websites []string = []string{
		"https://google.com",
		"https://fb.com",
		"https://instagram.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1) // The value sets the number of goroutine we have in this case we have only one goroutine so i set it to 1.
	}

	wg.Wait()
}

// greeter prints the given string 5 times with a small delay in between.
// This simulates some concurrent task.
// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		// Sleep for 3 milliseconds to simulate a time-consuming task.
// 		// This also helps demonstrate how goroutines allow other tasks to run concurrently.
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s) // Print the string passed to the function.
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	web, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error occured while acquiring the status code.")
		return
	}

	mut.Lock() // Locking the memory space by the goroutine.
	fmt.Printf("Status code for site %s is %v\n", endpoint, web.StatusCode)
	mut.Unlock() // Unlocking is required so that it can't go in the deadlock situation.
}