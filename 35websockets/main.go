package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade HTTP connections to websocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a websocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading", err)
		return
	}

	defer conn.Close() // It's crucial.

	// Listening for incoming messages.
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message", err)
			http.Error(w, "Error reading message "+ err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Recieved message: %s\\n", message)

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Failed to write message", err)
			http.Error(w, "Failed to write message "+err.Error(), http.StatusInternalServerError)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Websocket server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start the server")
	}
}