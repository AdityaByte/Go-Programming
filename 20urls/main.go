package main

import (
	"fmt"
	"net/url"
)

// Learning about fetching information from the url and constructing a url.
// There is another package inside the net library call as URL by which we can do these tasks.
// Url
const myUrl string = "https://www.google.com:2000/golang?courses=java&duration=4months"

func main() {
	fmt.Println("Url Handling in golang")

	// Here we have a package by which we can grab information about a url.
	parseUrl, _ := url.Parse(myUrl)

	fmt.Println(parseUrl.Scheme)   // Prints out the schema this which protocol we are using
	fmt.Println(parseUrl.Host)     // Prints out the host name with port (domain name)
	fmt.Println(parseUrl.Path)     // Path - endpoint.
	fmt.Println(parseUrl.RawQuery) // Prints out the query...

	queryParams := parseUrl.Query()
	fmt.Printf("Type of Query: %T\n", queryParams) // url.Values it is type of map...
	fmt.Println(queryParams["courses"])            // Prints out java.

	// Iterating through all the params
	for key, value := range queryParams {
		fmt.Printf("%v : %v\n", key, value)
	}

	// If you want to construct a url when we have paramerts
	// There we have a method url.URL for that but we have to pass a reference of it

	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "aditya.go.dev",
		Path: "golang",
		RawQuery: "devName=aditya",
	}

	finalURL := partsofUrl.String()
	fmt.Println(finalURL) // https://aditya.go.dev/golang?devName=aditya

}
