package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza store")
	fmt.Println("Give rating to our pizza in between 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//var input , _ = reader.ReadString('\n') // This is also valid.
	/*
	var input string
	var err string
	input, err = reader.ReadString('\n') // This is also valid.
	*/

	fmt.Println(input)
	fmt.Printf("The type of rating is %T \n", input)
	// Here we want to do that we want to add one in the input and show up on terminal.
	// But the user input is in the String format we have to convert that to integer.
	// For that we have a package called strconv for conversion
	// Note - when we press enter we get two things "{Rating}\n" so we have to remove the leading space from it using the strings package.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // It has two parameter one for the string that has to convert and the other for the size.

	// Here we have to handle the error
	if err != nil {
		fmt.Println(err)
		//panic(err) // Also do that this will show up the error and the program will end.
	} else{
		fmt.Println("The rating user gave us :", numRating + 1) // we have to remove the trailing space afterward of the string 
	}
}