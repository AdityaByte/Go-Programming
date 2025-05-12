package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json: "userna,e"`
	Password string `json: "password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request type: accept post request only")
		return
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Printf("The user request value %v\n", user)

	if user.Username == "Aditya" && user.Password == "Pawar" {
		tokenString, err := createToken(user.Username)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid credentials")
		return
	}

}