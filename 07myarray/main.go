package main

import "fmt"

// Another data type in Go is an array.
// An array is a collection of elements of the same type, and its size is fixed at compile time, unlike lists in other languages.
// Arrays are not commonly used in Go because slices provide more flexibility and are the preferred way to work with collections.

func main() {
    fmt.Println("Arrays in Go")

	// This is how we can create an array in go.
	var fruitList [4]string // declaration
	fruitList[0] = "mango"
	fruitList[1] = "jackfruit" // haha since it is not a fruit.
	fruitList[3] = "sitafal"

	fmt.Println("FruitList is :", fruitList) // It is of size 4 so it will create a space for the 2 index cause it is empty and we don't initialize it.
	fmt.Println("Size of fruitlist is :", len(fruitList))

	var vegList = [3]string{"mushroom", "kaddu", "bhindi"}
	fmt.Println("Vegy list is :", vegList)
	fmt.Println("Size of vegylist is :", len(vegList))

	// Go don't provide direct methods for sorting array and such common stuff that we can do in other langauges.
}
