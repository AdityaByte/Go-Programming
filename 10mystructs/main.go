package main

import "fmt"

// Struct means structure so we use them for creating a structure in go.
// Go follows functional programming paradigrm so it has no concept of class and class related things.
func main() {
	fmt.Println("Structs in golang")

	aditya := User{"aditya", 20, "aditya@go.dev", true} // This is how we use the structure

	// If we want to print out the structure.
	fmt.Println(aditya)
	fmt.Printf("The User Details are : %+v\n", aditya) // %+v placeholder prints out the user details more precisely.

	// We can also print seperate details of a user
	fmt.Printf("Name : %v and Email : %v\n", aditya.Name, aditya.Email)
}

// This is how we can create a struct in golang
// Syntax: type NameOfStructure struct { -- body -- }
// Everything in go is a type that's why.
// We use the first letter capital so that i can export it to different files because it is public.

// It's automatically indented the things because of intelligence - apart from knowledge.
type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
