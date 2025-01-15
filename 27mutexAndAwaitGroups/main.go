package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Races in golang")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	data := []int{0}

	// Here we have some anonymous function.
	// func(Parameters){Body}(Function call)
	// Here we have created a go routine.
	wg.Add(3) // No of goroutines
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("Function 1")
		mut.Lock()
		data = append(data, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("Function 2")
		mut.Lock()
		data = append(data, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("Function 3")
		mut.Lock()
		data = append(data, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(data)
}
