package main

import (
	"fmt"
	"net/http"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request type: accept get request only")
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	fmt.Println("tokenString:", tokenString)

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Missing authorization header")
		return
	}

	tokenString = tokenString[len("Bearer "):]

	fmt.Println(tokenString)

	err := verifyToken(tokenString)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid token")
		return
	}

	fmt.Fprintf(w, "Welcome to the protected area")
}
