package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Learning about base64 encoding.
// Base64 : Base64 is the binary-to-text encoding scheme that transforms the binary data into a sequence of printable ASCII characters, limited to a set of 64 unique characters.
// In golang there is a package which is prior to base64 is encoding/base64.

func main() {
	// Here we have to create a NewEncoder by the help of which we can encode the data.
	// It takes two things first is the encoding like hexencoding, stdencoding etc and second is the writer in which data it has to write.
	// As the another parameter is io.Writer so we want something that implements the io.Writer interface{}.

	var encodedData bytes.Buffer // We can write into the bytes.
	encoder := base64.NewEncoder(base64.StdEncoding, &encodedData)

	// Here we have to take the input from the user.
	fmt.Print("DATA: ")

	reader := bufio.NewReader(os.Stdin)
	actualData, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	if strings.TrimSpace(actualData) == "" {
		fmt.Println("ERROR: Null Data")
		return
	}

	encoder.Write([]byte(actualData))
	encoder.Close() // Close it here don't defer to close.

	encodedString := encodedData.String()

	fmt.Println("ENCODED STRING:", encodedString)

	// So now if we wants to decode the data this is how we can decode the data.
	// We have to setup a decoder here
	// Setup for reading the bytes of the encodedData
	// bytes := bytes.NewReader([]byte(encodedString)) // or we can also do this encodedData.bytes()
	bytes := bytes.NewReader(encodedData.Bytes())

	decoder := base64.NewDecoder(base64.StdEncoding, bytes) // It takes two things firstly in which encoding format we have to decode and secondly the reader where we have to read the data.

	decodedData, err := io.ReadAll(decoder)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	decodedString := string(decodedData)

	fmt.Println("DECODED STRING:", decodedString)

}
