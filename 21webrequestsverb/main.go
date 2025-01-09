package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Requests Verb in golang")
	// PerformGetRequest()
	PerformJsonPostRequest()
}

const url string = "http://localhost:8080/api"

func PerformGetRequest() {

	response, err := http.Get(url + "/get")

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

func PerformJsonPostRequest() {

	// This is not the genuine way of creating json but it is one of the way.
	requestBody := strings.NewReader(`
		{
			"id": 1234,
			"username": "aditya",
			"website": "aditya.go.dev"
		}
	`)

	// Here we are creating the Post request to the javaAPI
	response, err := http.Post(url+"/post", "application/json", requestBody)

	
	if err != nil {
		panic(err)
	}
	
	// Make sure you do that it's your responsibility.
	defer response.Body.Close()
	
	// Reading the content from the body
	content, _ := io.ReadAll(response.Body)

	// strings.Builder is a struct in the strings package
	// Here we are creating an instance of the Builder struct.
	// Builder acts as a mutable buffer where you can write content to make a string rather than creating multiple copies.

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is", byteCount)

	// Since the Builder is a mutable object we can return the instance here.
	fmt.Println(responseString.String())
}