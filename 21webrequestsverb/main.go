package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Requests Verb in golang")
	HandlingGetRequest()
}

func HandlingGetRequest() {
	const url string = "http://localhost:8080/api/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// we have to close the response body
	defer response.Body.Close()

	// Status code of response
	fmt.Println("Status code of response:", response.StatusCode)
	fmt.Println("Content length of response:", response.ContentLength)

	// Here we have to read the response

	content, _ := io.ReadAll(response.Body)
	//fmt.Println(content) // Raw content
	// fmt.Println(string(content))

	// There is an another way of converting or reading the response 
	// By strings package

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content) // This will take the content and write it.

	fmt.Println("byte count is", byteCount)
	fmt.Println(responseString.String())
}
