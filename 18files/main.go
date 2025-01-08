package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"os"
)

// Learning about some basic file operation like-reading,writing,creation etc.
func main() {
	fmt.Println("Files in golang")

	// Content
	content := "Hey i am writing in the file. \nUsing golang code."

	// File creation is a OS operation.
	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	//Panic is the function which takes error and stops the program and throw out the error.
	// 	panic(err)
	// }
	checkNilErr(err)
	// Using the WriteString function we can write the string in file.
	// It takes up two arguments one the file and the content
	// length, err := file.WriteString(content)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("The length of the file content is:", length)
	// Now we need to close the file cause it has many function open
	defer file.Close()

	// If we want to read the file in golang this is how we can read
	// File creation is a os operation but reading a file is not an os operation.
	// For that we have a library in go called as ioutil
	// But since now the ioutil.ReadFile is deprecated now we can read the file using the os.
	// dataByte, err := ioutil.ReadFile("./mygofile.txt") // This will return the data in the form of bytes
	dataByte, err := os.ReadFile("./mygofile.txt")
	checkNilErr(err)
	fmt.Println("File data is :", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
