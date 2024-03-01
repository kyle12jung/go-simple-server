package main

import (
	"encoding/json"
	"fmt"
	"go-practice/todo"
	"net/http"
)

type Message struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Message{Status: 200, Data: "Default endpoint"}
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/todos", todosHandler)
	fmt.Print("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		todo.HandleCreate(w, r)
	case "GET":
		todo.HandleList(w, r)
	case "PUT":
		todo.HandleUpdate(w, r)
	case "DELETE":
		todo.HandleDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Message{Status: http.StatusMethodNotAllowed, Data: "Method Not Allowed"})
	}
}
