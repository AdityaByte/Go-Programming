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
	// EncodingJson()
	DecodeJson()
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

func DecodeJson() {
	// The data we get is in the byte format further ado we make operations on it.
	jsonDataFromWeb := []byte(`
		{
			"carname": "lamberghini aventador",
			"company": "lamberghini",
			"tags": [
					"Fastest",
					"luxory"
			],
			"ownerName": "Aditya Pawar"
		}
	`)

	// Firstly we have to check the json is valid or not.
	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		// Decoding the json data using unmarshal method
		// Decoding with the help of struct reference.
		var car1 car                           // We just declare it doesn't initialize it.
		json.Unmarshal(jsonDataFromWeb, &car1) // Giving the reference of the car1.
		fmt.Printf("%#v\n", car1)             // Placeholder %#v - format specifier.
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// If we want to decode json in just key value pair then...

	var myCarDetail map[string]interface{} // since we dont know about the type of value which is in the value so we use the interface{}

	json.Unmarshal(jsonDataFromWeb, &myCarDetail)

	fmt.Printf("%#v\n", myCarDetail)

	// Can also consume that using the for loop

	for k,v := range myCarDetail {
		fmt.Printf("Key : %v and Value : %v and Type : %T\n", k, v, v)
	}
}
