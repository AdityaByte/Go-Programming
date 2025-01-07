package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// If else comes in the control flow statement because they control the flow of the program.
func main() {
	fmt.Println("If Else statement in golang")
	reader := bufio.NewReader(os.Stdin)

	// Taking age as input
	// input, err := reader.ReadByte() // Cannot use this function cause it reads a single byte.

	fmt.Printf("Age : ")
	input, _ := reader.ReadString('\n')

	// Here we are converting the input to integer
	// The ParseInt function converts out the bit size to 64 for integer.
	age, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32) // Second parameter wants base so it would be 10 for decimal.

	if err != nil {
		panic("Getting some error while converting string to integer")
	}

	fmt.Printf("Type of input %T\n", age)

	// That's the syntax of if else statement in golang and you need to do this as same because the lexer checks out everything that
	// You have written syntetically correct or not...
	if age < 18 {
		fmt.Println("Bro, You can't drive...")
	} else if age >= 18 && age < 50 {
		fmt.Println("Bro You can drive anything car bus aeroplane...")
	} else {
		fmt.Println("You can't drive dude sorry...")
	}

	// There is an another syntax in go by which you can initialize a variable too..

	if num := 34; num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
