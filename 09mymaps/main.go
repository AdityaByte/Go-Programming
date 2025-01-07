package main

import "fmt"

// Maps are the hashtables and you can compare them to dictionary in python or map in set
// which stores data in a key vay pair is called a map.
// A map can allow duplicate.
func main() {
	fmt.Println("Maps in go")
	// This is how we can create a map in go
	// Syntax: make(map[typeOfKey]typeOfValue)
	// Note - You can also define the size that how many values you're going to store.
	languages := make(map[string]string)
	languages["GO"] = "GoLang"
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["HS"] = "Haskell"

	fmt.Println("List of languages :", languages)

	// If we want to delete an item from the map then
	delete(languages, "JS") // 1st argument - Map, 2nd argument - key.

	fmt.Println("New Map :", languages)

	// Looping through a map (Since we haven't studied looping statement in go)

	for key, value := range languages {
		fmt.Printf("Key %v is short form of %v \n", key, value) // %v is a placeholder for value.
	}
	// Note - Here you can also use the comma ok syntax with the walrus operator if you want only key or only value.

	var courseRating = map[string]int{
		"Golang":     5,
		"Python":     4,
		"Java":       5,
		"Javascript": 3,
	}

	fmt.Printf("%v \n", courseRating)

}
