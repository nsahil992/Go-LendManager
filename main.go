package main

import (
	"encoding/json" // Importing the JSON package for encoding and decoding
	"fmt"          // Importing the fmt package for formatted I/O
	"net/http"     // Importing the net/http package for HTTP server and client
)

// lent holds the mapping of friends to items lent to them
var lent = map[string][]string{
	"Courage": {},
	"Ben":     {"Watch"},
	"Mr.Bean": {"Teddy"},
}

func main() {
	// Setting up the HTTP handlers for different API endpoints
	http.HandleFunc("/api/takeback", handleTakeback) // Handler for taking back an item
	http.HandleFunc("/api/give", handleGive)         // Handler for lending an item
	http.HandleFunc("/api/newfriend", handleNewFriend) // Handler for adding a new friend

	// Starting the server on port 8080
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil) // This line listens for incoming HTTP requests
}

// handleTakeback handles the request to take back an item from a friend
func handleTakeback(w http.ResponseWriter, r *http.Request) {
	friend := r.URL.Query().Get("friend") // Getting the friend's name from the query parameters
	if items, ok := lent[friend]; ok && len(items) > 0 { // Check if the friend exists and has lent items
		lent[friend] = lent[friend][:len(items)-1] // Remove the last item lent to the friend
		response := map[string]string{"message": "Item taken back from " + friend} // Prepare the response
		json.NewEncoder(w).Encode(response) // Encode the response as JSON and send it
	} else {
		http.Error(w, "No items found for this friend", http.StatusNotFound) // Return 404 if no items found
	}
}

// handleGive handles the request to lend an item to a friend
func handleGive(w http.ResponseWriter, r *http.Request) {
	friend := r.URL.Query().Get("friend") // Getting the friend's name from the query parameters
	item := r.URL.Query().Get("item") // Getting the item name from the query parameters
	lent[friend] = append(lent[friend], item) // Append the item to the friend's list
	response := map[string]string{"message": "Lent " + item + " to " + friend} // Prepare the response
	json.NewEncoder(w).Encode(response) // Encode the response as JSON and send it
}

// handleNewFriend handles the request to add a new friend
func handleNewFriend(w http.ResponseWriter, r *http.Request) {
	friend := r.URL.Query().Get("friend") // Getting the friend's name from the query parameters
	if _, exists := lent[friend]; !exists { // Check if the friend already exists
		lent[friend] = []string{} // Initialize the friend's list with an empty array
		response := map[string]string{"message": friend + " added as a new friend"} // Prepare the response
		json.NewEncoder(w).Encode(response) // Encode the response as JSON and send it
	} else {
		http.Error(w, "Friend already exists", http.StatusConflict) // Return 409 if friend already exists
	}
}
