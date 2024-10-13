package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var lent = map[string][]string{
	"Courage": {},
	"Ben":     {"Watch"},
	"Mr.Bean": {"Teddy"},
}

func startServer() {
	http.HandleFunc("/api/takeback", handleTakeback)
	http.HandleFunc("/api/give", handleGive)
	http.HandleFunc("/api/newfriend", handleNewFriend)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func handleTakeback(w http.ResponseWriter, r *http.Request) {
	friend := r.URL.Query().Get("friend")
	if items, ok := lent[friend]; ok && len(items) > 0 {
		lent[friend] = lent[friend][:len(items)-1]
		response := map[string]string{"message": "Item taken back from " + friend}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "No items found for this friend", http.StatusNotFound)
	}
}

func handleGive(w http.ResponseWriter, r *http.Request) {
	friend := r.URL.Query().Get("friend")
	item := r.URL.Query().Get("item")
	lent[friend] = append(lent[friend], item)
	response := map[string]string{"message": "Lent " + item + " to " + friend}
	json.NewEncoder(w).Encode(response)
}

func handleNewFriend(w http.ResponseWriter, r *http.Request) {
	friend := r.URL.Query().Get("friend")
	if _, exists := lent[friend]; !exists {
		lent[friend] = []string{}
		response := map[string]string{"message": friend + " added as a new friend"}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Friend already exists", http.StatusConflict)
	}
}
