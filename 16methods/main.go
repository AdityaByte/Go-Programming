package main

import "fmt"

// Methods : Methods are similar to functions but they are associated with some types.

func main() {
	aditya := user{"aditya", 20}
	fmt.Println(aditya)              // {aditya 20}
	fmt.Println(aditya.age, aditya.name)

	// Let we have to change the name for aditya and age too
	aditya.setData("Aditya Pawar", 21)
	fmt.Println(aditya) // Now the fields are changed..

	aditya = aditya.setName("Somebody") // If we reinitialize the aditya variable then the name will be changed
	fmt.Println(aditya) // It will change because setName returns a new copy of user
}

// The struct and its fields are private to the package because they start with lowercase letters.
// These private fields can be accessed within the same package, but if the struct is in another package,
// we cannot access the fields because they are private (lowercase letters).
type user struct {
	name string
	age  int
}

// Declaring a method:
// Methods in Go are functions associated with a specific type (usually structs).
// The method must have a receiver, which allows it to operate on an instance of the type.
// You can declare methods on any type, such as structs or interfaces, but it’s most common with structs.
// func () nameOfMethod(Parameter) returnType {}

// Here we are using a pointer to the struct because we want to modify the original struct.
// This method is similar to a setter, and since it uses a pointer receiver, 
// the changes to the fields will persist.
func (u *user) setData(name string, age int) {
	u.name = name
	u.age = age
}

// Without pointer, changing the name
// This method uses a value receiver, so it returns a new copy of the struct with the updated name.
// To see the updated name, you must reassign the variable.
func (u user) setName(newName string) user {
	u.name = newName
	return u
}
