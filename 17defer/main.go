package main

import "fmt"

// Defer in english means - to delay something until a later time; or to postpone;
// If you place the defer keyword to a statement that means the execution is held at last.
// Defer is something like stack
/* Same like stack follows the LIFO Operation of last in fast out so the defer statement are 
pushed into a stack
example : for i := 1; i <= 5; i++ {
			defer fmt.Println(i)  
		}
// normal loop - 1 2 3 4 5
but the 1 goes to the stack then 2 and so on
// Stack 
1 2 3 4 5 -> open
print : 5 4 3 2 1
*/
func main() {
	// Function likely read or execute as top to bottom or left to right
	defer fmt.Println("Hey this is defer") // This statement is going to execute at last
	fmt.Println("Defer in golang")
	mydefer()

	// Output
	/*Defer in golang
	stack -> mydefer() [5,4,3,2,1] then "Hey this is defer"
	*/
}



// let us take an example that we have a function......
func mydefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i) // 
	}
}