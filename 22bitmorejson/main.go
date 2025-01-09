package main

import (
	"encoding/json"
	"fmt"
)

// Note - It's a good practice to have a json key in lowercase format.
type car struct {
	Name        string   `json:"carname"`
	Company     string   `json:"company"`
	NumberPlate int      `json:"-"` // I want this is not visible for not sended.
	Tag         []string `json:"tags"`
	Owner       string   `json:"ownerName,omitempty"` // omit empty omits the empty field and don't print that.
}

func main() {
	fmt.Println("Working more with json")
	EncodingJson()
}

// Encoding data into json
func EncodingJson() {
	// For encoding the struct car into json
	// Golang has a package called encoding/json
	myCars := []car{
		{"Honda city", "Honda", 1111, []string{"Fast", "sleek"}, "Aditya"},
		{"lamberghini aventador", "lamberghini", 0101, []string{"Fastest", "luxory"}, "Aditya Pawar"},
		{"Audi cls", "volkswagon", 2334, []string{"Good"}, ""},
	}

	// We have to encode the myCars data to json

	// encodedJsonData, err := json.Marshal(myCars) // In ugly format
	encodedJsonData, err := json.MarshalIndent(myCars, "", "\t") //It expects three things type,prefix and indent

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", encodedJsonData) // we print json by placeholder %s
}
