package main

import (
	"bufio"
	"fmt"
	"os"
)

// Here let us we are creating for taking a user input.
func main() {
	
	welcome := "Welcome Here we take user input" // Here we are using walrus operator so that go runtime can detect the type based on value.
	fmt.Println(welcome)
	fmt.Println("What do you want to rate our pizza : ")
	// For taking out the input from user we have a library known bufio
	reader := bufio.NewReader(os.Stdin) // Setting up the reader where we declare that we have to take input from the os and its a standard input.

	// In golang we don't have try catch block but we have
	// Comma Ok syntax | Comma Error syntax
	// The underscore variable takes the error.
	// Here we use _ cause we only want the input not the error that might happen.
	input, _ := reader.ReadString('\n') // Here we have to specify from where we have to read string so we say that when the line ends by \n.

	// That's the comma ok syntax 
	_, err := reader.ReadByte()
	fmt.Println(err) // we get nil cause no error was caught.

	fmt.Println("The user gives rating of", input)
	fmt.Printf("The type of the user rating is %T", input)

	// There is another way of reading input till the white space character ' '.
	// via using the fmt.Scanln() function or fmt.Scan() function but they only reads till the first space.

	var data string
	fmt.Scan(&data)

}	