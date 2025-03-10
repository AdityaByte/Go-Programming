package main

import (
	"fmt"
	"sync"
	"time"
)

// Here we are creating out two methods one for reading and one for writing
// The normal mutex (mutual exclusion lock) does not distinguish between the read locks and write locks.
// We need rw-mutex when we writing a data we need lock but when we are reading it why we are blocking the go routine 
// so in this situation we use the read write mutex.

var count int

var rwMutex sync.RWMutex

func readTask(name string) {
	rwMutex.RLock() // This is the read lock
	fmt.Println(name, "is reading the value of count:", count)
	time.Sleep(time.Millisecond * 400)
	rwMutex.RUnlock() // This is the read unlock
}

func writeTask(name string) {
	rwMutex.Lock() // This is the normal lock
	fmt.Println("Increasing the value of count...")
	count++
	time.Sleep(time.Millisecond * 400)
	rwMutex.Unlock()
}

func main() {
	wg := sync.WaitGroup{} // We have created an instance of the waitgroup here.

	wg.Add(3)

	go func(name string) {
		for i := 0; i < 10; i++ {
			writeTask(name)
		}
		wg.Done()
	}("Thread-write")

	go func(name string) {
		for i := 0; i < 10; i++ {
			readTask(name)
		}
		wg.Done()
	}("Thread-read1")

	go func(name string) {
		for i := 0; i < 10; i++ {
			readTask(name)
		}
		wg.Done()
	}("Thread-read2")

	wg.Wait()
	fmt.Println("All task done")
}
