package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning about the time package in golang")

	presentDate := time.Now(); // It gives out the present time
	fmt.Println(presentDate) // It is in a prettier uglier format we can format the time.
	// For formatting out the time....
	/* Things you have to remeber before formating
	* you can write this can be 01-02-2006 or 01/02/2006 or 02/01/2006 here 02 denotes day 01 denotes month 
	* day can be anything like 02,03,04 but month format must be 01
	* For day you've to write Monday. Note - It is case sensitive.
	* for time formatting this is the format that 15:04:05 15-hour 04-minutes and 05-seconds.
	*/
	fmt.Println(presentDate.Format("02/01/2006 Monday 15:04:05")) // we have to pass out a string which must be the same as per the docs.

	// If we want to create our own date then this is how we have to do.
	createdDate := time.Date(2024, time.December, 31, 23, 59, 59, 59, time.Now().Location())
	fmt.Println(createdDate) 
	// We can format it too
	// These format is in 24 hour clock format
	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
	// If we want to format that in 12 hour clock format then..
	fmt.Println(createdDate.Format("01-02-2006 Monday 03:04:05 PM"))
}
