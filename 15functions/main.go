package main

import "fmt"

// Main function
// So it is the entry point of the program
// We are not calling this function everywhere but the execution starts from there only.
/*
Functions: A function is an independent block of code that performs a specific task. 
It is not associated with any specific type.
*/
func main() {
	fmt.Println("Functions in golang")

	// Normal adder function
	result := adder(10, 20)
	fmt.Println("Sum of two numbers is :", result)

	proResult, message := proAdder(1, 2, 3, 4, 5, 6, 7) // If you want to ignore the message you can place underscore there.
	fmt.Println("Sum of many numbers are :", proResult)
	fmt.Println("Message :", message)

	var newString string = concatingString("aditya ", "pawar")
	fmt.Println("New string is :", newString)

	// Here we have anonymous function too that don't have a name 
	// You can call anonymous function immediately.
	func (message string) {
		fmt.Println("Message is", message)
	}("Hello world")
}

// Function in golang
/*
- Functions are the block of code which are used for performing a particular operation
- There can be multiple functions in a program which do multiple tasks.
- Syntax:
func nameOfFunction(Parameters) returnType {
	// Body of the function
}
// Note - The returnType is known as the signature of the method in golang.
*/
// Note - If the first letter is small then the function is not exported in another file
// Cause  it is private and if it is Capital then it is public.
func adder(value1 int, value2 int) int {
	return value1 + value2
}

// Varargs in golang : Variadic in golang 
// If we don't know how many arguments are passing by the user
// The comma ok syntax is same based on function like this here we are returning
// two things first the value and secondly the message.
func proAdder(numbers ...int) (int, string) {
	result := 0
	for _, value := range numbers {
		result += value
	}

	return result, "Hey this is the pro adder"
}

// Note - In the function parameter you have to define which type of arguments you are expecting.
// Other wise it will give an error

// If multiple paramerts have the same type you can do this shortcut

func concatingString(string1, string2 string) string {
	return string1 + string2
}
