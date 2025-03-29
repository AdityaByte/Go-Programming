package main

import (
	"context"
	"fmt"
	"time"
)

// Let us take an example for understanding the concept of context api.
// Suppose u go to a restaurant and u don't want to wait forever so u tell the waiter
// if the order doesn't came in 10 minutes cancel my order.

// Same goes with the context api it is being used for managing timeouts, deadlines and cancellations
// across various parts of the program.

// Context is always passed as pass-as-argument.

func anotherFunc() (data int) {
	time.Sleep(time.Millisecond * 100)
	return 100
}

func simpleExample() {

	// Creating a context with timeout.

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)

	defer cancel() // Necessary to close the resource.

	// Passing context as argument to a function

	// Creating a channel which accept the data
	result := make(chan int)

	go func(){
		result <- anotherFunc() // It triggers a slow process which takes 400 miliseconds but our context only waits for 200 miliseconds
	}()

	// For select loop
	select {
	case res := <-result:
		// fmt.Println("Result is:", result) //Basically it is printing the reference
		fmt.Println("Result is:", res)
	case <-ctx.Done():
		fmt.Println("Request is taking too long to process.")
	}

}
