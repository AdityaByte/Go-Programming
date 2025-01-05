package main

import "fmt"

//jwtToken := 2391121 // This is not allowed
// But these are allowed
// var jwtToken int = 231312
// var jwtToken2 = "Hello world"

// Constants - We cannot change the value of the constant since we can change the value of a variable.

// Note - Here i have created the const with the first letter capital
// Capital L here the significant importance so now this is the public variable
// In Java it is equivalent to public static String loginToken;
const LoginToken string = "loginToken231"

// Main function defination - Entry point of the go program.
// Here we are studying how do we declare and initialize a variable in go.
func main() {

	// Syntax var(keyword) variableName type - declartion = initialization
	var username string = "Aditya"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n" , username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable if of type : %T \n", isLoggedIn)

	// In case of integer we have two types unsigned and signed integer
	/*
	1. Unsigned - unsigned are those who can hold positive values only
	In golang we refer to them as uint,uint8,uint16,uint32,uint64 
	Here the number denotes that how much bits we can hold in them 8 bit, 16 bit etc.

	2. Signed - Signed are those who can hold both the values positive and negative both
	In golang we refer to them as int,int8,int16,int32,int64
	Here the number denotes that how much bits we can hold in them 8 bit, 16 bit etc.
	*/

	var number uint8 = 255 // If we convert that to 256 it can't hold it and gives an error cause it is out of range
	fmt.Println(number)
	fmt.Printf("Variable is of type : %T \n", number)

	// Floating point values 
	// In go we have float32 for 32 bits and float64 for 64 bits value.
	var decimalValue float32 = 1212.12128782738223 // If we have return such more things after decimal but it can't take it cause it is a float32
	fmt.Println(decimalValue)
	fmt.Printf("Variable is of type : %T \n", decimalValue) // Prints - 1212.1213 rounds that

	var decimalValue2 float64 = 1212.12128782738223
	fmt.Println(decimalValue2)
	fmt.Printf("Variable is of type : %T \n", decimalValue2)

	// There we have another Type known as complex for dealing with complex numbers imaginary values
	// complex64 with float32 real and imaginary parts
	// complex128 with float64 real and imaginary parts

	var complexNumber complex64 = complex(2, 4)
	fmt.Println(complexNumber)
	fmt.Printf("Variable is of type : %T \n", complexNumber)


	// Here we have an alias for int8 which is byte
	// Here we have byte too which is equivalent to uint8 of 1 byte 0 - 128
	var byteValue byte = 128
	fmt.Println(byteValue)
	fmt.Printf("Variable is of type : %T \n", byteValue)

	// If we don't have the idea of how much memory we can use you can explictly define it as int.
	var intValue int = 2121;
	fmt.Println(intValue)
	fmt.Printf("Variable is of type : %T \n", intValue)

	// There we have an another alias for int32 is rune
	var int32Value rune = 12121212
	fmt.Println(int32Value)
	fmt.Printf("Variable is of type : %T \n", int32Value)

	// Default values of some alias
	var anotherVariable int // here we just declare the variable not initialize it.
	fmt.Println(anotherVariable) // Zero is the default value for integers.
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// Implicit way of declaring a variable
	// Here the lexer comes into the game if you are not defining the Type for the variable 
	// then the lexer would decide it for you based on the value which you provided.
	var websiteName = "golang.org"
	fmt.Println(websiteName)
	fmt.Printf("Variable is of type : %T \n", websiteName)
	
	// no var style
	// := short variable declaration operator, walrus operator
	// But we can only use this opertor inside the function not outside it.
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type : %T \n", numberOfUsers)

	// Using the public variable LoginToken
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}