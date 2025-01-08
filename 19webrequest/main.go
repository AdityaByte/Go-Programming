package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://127.0.0.1:5000/pyapi"

func main() {
	fmt.Println("Web Request in golang")
	// Get method reside in net/http library through which we can send the get request.
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// ioutil.ReadAll(response.Body) // its deprecated now.
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response: %T\n", databytes) // This will print out the message of the python api
	fmt.Println("Response :", string(databytes))
}
