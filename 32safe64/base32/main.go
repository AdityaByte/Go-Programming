package main

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"fmt"
	"io"
	"log"
	"os"
)

// First of all read the base64 folder main.go ok where i have mentional more clearly
// here i am just doing some coding stuff like how we can encode 32bit data ok.

func main() {

	fmt.Print("DATA: ")
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err);
	}

	var encodedData bytes.Buffer
	encoder := base32.NewEncoder(base32.HexEncoding, &encodedData)
	encoder.Write([]byte(data)) // Although it returns the size and the error you can catch it. (You have to catch it)
	encoder.Close()
	fmt.Println("Encoded Data:", encodedData.String())

	// If we want to decode the data.
	decoder := base32.NewDecoder(base32.HexEncoding, bytes.NewReader(encodedData.Bytes()))
	decodedData, err := io.ReadAll(decoder)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decoded Data:", string(decodedData))
}