package main

import "fmt"

// The pointers in golang is same as C and C++, they are simple at the same time they are interesting too
// Pointers are the variable that stores address.
/*
When you want to avoid unnecessary memory allocation and ensure that the actual value 
(not a copy) is passed, you should use pointers. A pointer is simply a variable that holds the
memory address of another variable, allowing direct manipulation of the original value without
creating additional copies.
*/

// Pointers are created with the astring(*) sign and pointer stores the address.
// Ampersand(&) is used as a reference or it prints out the address.
func main() {
	fmt.Println("Introduction to pointers")

	// Here we have created a variable which is holding the value 34
	number := 34
	
	// This is how we can create a pointer
	var ptr *int = &number
	// This is also how we can create a pointer.
	var newPtr = &number

	fmt.Println(ptr, newPtr) // This will print out the reference to the number variable.
	fmt.Println(*ptr) // Here, we are dereferencing the pointer to access the actual value it points to.

	*ptr = *ptr + 1 // Here we are derefercing the pointer and increasing the actual value by 1.
	fmt.Println("The actual value of number is now ", number) // Prints out 35


	// Double pointer - Double pointer stores the address of itself the pointer

	var doublePtr **int = &newPtr
	doublePtr2 := &ptr
	fmt.Println(doublePtr, doublePtr2) // It gives out the address of the both pointers.

}