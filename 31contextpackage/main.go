package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*Context is like a bag or container that holds information that is shared between
different parts of the program, especially when it comes to handling a request.
This information can include things like timeouts, cancellation signals,
and other data that is specific to that request.

For example, imagine you are building a web server that handles a lot of incoming requests.
Each request has its own specific needs and requirements,
such as a deadline for how long it should take to complete.
The context allows you to keep track of these individual requirements for each request,
and make sure that they are handled properly. */

/* # Context Package :
-> Let us understand this package with some use case
-> Let we have a function which do some third party api stuff like fetching n all.
-> So we don't know how much time the third party server is going to take to serve the request.
-> So to define a particular time by which if the request data didn't serve on we will close the context.
-> This can be done using the context package in golang.
*/

func main() {
	fmt.Println("Context package in golang")

	// Here we are calculating the time like how much it is gonna take
	start := time.Now()

	// Here we have to create a context at background
	// ctx := context.Background()

	// Creating a context with key value
	ctx := context.WithValue(context.Background(), "aditya", "pawar")
	val, err := fetchUserData(ctx, 100)

	if err != nil {
		log.Fatal("Error occured:", err)
	}

	fmt.Println("Result:", val)
	fmt.Println("Time taken:", time.Since(start))

	simpleExample()
}

type Response struct {
	result int
	err    error
}

// Here this is the function which calls the fetchThirdPartyStuff() function and it too creates a context.
func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200) // Here we created an another context which shutsdown if the request takes more than 200 miliseconds.
	defer cancel()

	responseChannel := make(chan Response)

	// Since we have created a goroutine for handling the api stuff so we cannot handle
	// the error efficiently inside it so we have to create a struct namely response.
	go func() {
		val, err := fetchThirdPartyStuff()
		responseChannel <- Response{
			result: val,
			err:    err,
		}
	}()

	for {
		select {
		// In the below case if 200 miliseconds has been triggerred hai the api stuff didn't fetch the response then
		// we got the empty value which is kinda similar to the ctx.Done() which is an empty struct.
		case <-ctx.Done():
			return 0, fmt.Errorf("Fetching data from third party takes too long.")
		case resp := <-responseChannel:
			return resp.result, resp.err
		}
	}
}

func fetchThirdPartyStuff() (int, error) {
	time.Sleep(time.Millisecond * 100) // Here let us assume that is request is taking 500 miliseconds to serve.
	return 666, nil
}
