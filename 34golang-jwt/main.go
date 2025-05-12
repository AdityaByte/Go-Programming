package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/protected", ProtectedHandler)

	fmt.Println("Starting the server")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server %s", err)
	}
}
