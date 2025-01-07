package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Math package in golang")

	// By the help of math package in golang we can perform math operations

	// Constant Value 
	piValue := math.Pi
	episilonValue := math.E
	logBase10 := math.Log10E

	fmt.Println(piValue, episilonValue, logBase10)
	
	// Some function example 

	// Let we have constant constants
	const (
		a = 121
		b = 146
	)
	fmt.Println(math.Sqrt(a)) // This will calculate the square root of a number
	fmt.Println(math.Sqrt(b))

	// If you want to calculate power
	fmt.Println(math.Pow(2, 4)) // 2^4 is 16

	// If you want to round off a number
	var num float32 = 1234.52345
	fmt.Println(math.Round(float64(num)))

	// Random package inside the math for random operations
	// Here we have a function Intn for generating random number for range n
	// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source. It panics if n <= 0.
	randNumber := rand.Intn(6) + 1
	fmt.Println("Random number :", randNumber)
}
