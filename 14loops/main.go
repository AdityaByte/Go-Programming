package main

import "fmt"

// In go there is only one loop which is for and using different condition we can make it functionable like while for each and all.
func main() {
	fmt.Println("Loops in golang")

	// Let here we have a slice of days
	days := []string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Slice of days:", days)

	// Here we want to iterate through this slice
	// In go there is only post-increment no pre-increment.
	fmt.Println("Another example of loop :")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// Another way we can define this
	// Just Like other languages for-each loop it doesn't give the exact value here it gives out the index
	// The reason for that is go is built differently.
	fmt.Println("Another example of loop :")
	for i := range days {
		fmt.Println(days[i])
	}

	// If you want the value too you can use the comma ok syntax for that
	fmt.Println("Another example with comma ok syntax :")
	for index, value := range days { // Can also write that range(days)
		fmt.Printf("%v : %v\n", index, value)
	}

	// While loop in go by for keyword
	fmt.Println("While loop example using for keyword : ")
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			rougueValue++
			continue // Don't write just continue it you have to define the case for which you have to skip
		}

		if rougueValue == 7 {
			goto aditya // goto statement in go
		}

		if rougueValue == 9 {
			break // Break the loop.
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

// Here we are defining a label
aditya:
	fmt.Println("Jumping out from the loop....")
}
