// Channels - Channels are the way by which our multiple goroutines can talk to each other.

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	// Channels are the pipeline by which your multiple can share out the data.
	myCh := make(chan int, 2) // This is the type of unbuffered channel with no capacity
	// If we want to push a value to a channel then this is how we can push a value to a channel
	// Note - The arrow always points to the left in golang.
	//myCh <- 5

	//fmt.Println(<-myCh)

	// In the upper scenerio we will get an error cause we will pass a value to a channel when a guy is listening to me.
	// So here the line which is listening to the channel is at 16 and we are passing a value at 14
	// for resolving the issue we can do this....

	//fmt.Println(<-myCh) // It still gives us an error.
	//myCh <- 5

	// Here we have to create a wait group and a goroutine for listening and sending out the values

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Here this anonymous method is going to recieve the channel data.
	// We can add up a signature -> ch <-chan int.
	// RECIEVE Only go routine.
	go func(ch <-chan int, wg *sync.WaitGroup){

		// Here is a way by which we can checkout that the channel is open or close
		// By default the channel gives up two values which is the value itself and a boolean value.

		value, isChannelOpened := <-ch

		fmt.Println(isChannelOpened)
		fmt.Println(value)

		// fmt.Println(<-ch)
		// fmt.Println(<-ch)
		wg.Done()
	}(myCh, &wg)

	// Here this method tells here we are sending data to the channel.
	// Signature for this channel since channel is a box - ch chan<- int
	// SEND Only go routine.
	go func(ch chan<- int, wg *sync.WaitGroup){

		// Let us talk about the situation when we are sending the values after the channel has been close then it will gives an error
		// Listening is not problematic if we are not sending any values and we are listening then it will gives 0.
		ch<-50
		close(ch)
		// ch<-10
		// ch<-20
		// You can close a channel too that the work has been done.
		// close(ch)
		wg.Done()
	}(myCh, &wg)

	wg.Wait()

	// Note - If you using a unbuffered channel and you are sending more values then listening to it then it will 
	// Through up an error so we can define a buffered channel with capacity.
}
