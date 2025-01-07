package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and Case in golang")

	dice := rand.Intn(6) + 1 // From this we get the random number.

	// You don't need to write break in each case just like languages like C,C++ and java python.
	// If you want to proceed to check more case further you can define fallthrough in that..
	switch dice {
	case 1: 
		fmt.Println("Oh nice, you got 1 you can start now..")
	case 2:
		fmt.Println("You can move to 2 place")
	case 3:
		fmt.Println("You can move to 3 place")
		fallthrough // now it checks for next case too and break out if fallthrough is not defined there.
	case 4:
		fmt.Println("You can move to 4 place")
	case 5:
		fmt.Println("You can move to 5 place")
	case 6:
		fmt.Println("You can move to 6 place")
	default:
		fmt.Println("What's this number?")
	}
}
