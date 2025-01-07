package main

import (
	"fmt"
	"sort"
)

// In Slices we are not gonna define how many values are come out there when we define
// how many values are coming then we are creating an array.
func main() {

	fmt.Println("Slices in go")

	// This is how we can create slices in java
	// As we are talking about slices so, slices are under the hood are arrays.
	// Once they get some abstraction layer and some more features they are called slices.

	var slicesVariable []int // If we don't define the size here this means we are creating a slice
	fmt.Println(slicesVariable)
	fmt.Println(len(slicesVariable))
	fmt.Printf("The type of slice variable is : %T \n", slicesVariable)

	// Since slices are dynamic.
	// We use the append method for adding a item in the slice
	slicesVariable = append(slicesVariable, 10, 11, 12)
	fmt.Println(slicesVariable)
	fmt.Println(len(slicesVariable))

	// There is an another way of creating slices too
	// If we are using this syntax then we have to initialize it too.
	//var slicesVariable2 = []string{}
	var fruitList = []string{"mango", "tomato"}
	fruitList = append(fruitList, "orange", "dragon-fruit")
	fmt.Println(fruitList)

	// we can also create that using the walrus operator too.
	vegyList := []string{
		"Jack-Fruit",
		"Singhada",
		"cherry", // If i am doing this we have to paste out a comma there.
	}
	fmt.Println(vegyList)

	// Slicing in golang in slices
	// The colon syntax is used for slicing a slice.
	vegyList = vegyList[:1] // 1st index - inclusive and last index - exclusive.
	fmt.Println(vegyList)

	// We have studied that in memory allocation that there are two ways for allocating space in memory is using new or make keyword.
	// Here in this case we are allocating a space for a datatype slice which can hold 4 values.
	highScores := make([]int , 4)

	highScores[0] = 234
	highScores[1] = 978
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777 // It gives an error

	/*
	So what's going on there is that 
	the make keyword creates a slice that can store the value of upto 4 integer but this 
	is the default location of the memory,
	but as soon as we use the method append it is going to re-allocate the memory
	and all of the values will accomodated.
	*/
	highScores = append(highScores, 555, 444, 666)

	fmt.Println(highScores)

	// There is a package sort this is used for sorting operation and other things too

	sort.Ints(highScores)
	fmt.Println("Sorted Slice :", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // Returns true

	// How to remove a value from the slice based on index.

	// Let we have a slice of cars
	var cars = []string{"hyundai", "toyota", "thar", "mercedes", "audi", "ferrari"}
	fmt.Println("Original car list :", cars)
	// For removing a item from a slice we have append method we can remove an item too using the append method
	var index int = 2
	cars = append(cars[:index], cars[index+1:]...) // We later study about the three dot.
	fmt.Println("New list of car :", cars)

	// Slice of type rune in go.
	charList := []rune{'a','b','c'}
	fmt.Printf("%c \n", charList)
	charList[0] += 1
	fmt.Printf("%c \n", charList)

	// Here if we want to remove two items from a slice 
	numbers := []int{1,2,3,4,5,6,7}
	// so i want to remove 3 and 6 from that.
	index1 := 2
	index2 := 5
	// Then
	numbers = append(numbers[:index1], numbers[index1+1:]...)
	fmt.Println(numbers)
	// Here we have removed one item so the index2 is decreased by 1
	numbers = append(numbers[:index2-1], numbers[index2:]...)
	fmt.Println("Final slice of numbers :", numbers)

}
