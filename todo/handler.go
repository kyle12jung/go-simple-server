package todo

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos := GetTodos()
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		return
	}
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	var todo TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo = CreateTodo(todo)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(todo)
	if err != nil {
		return
	}
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var todo TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedTodo, updatedStatus := UpdateTodo(todo)
	if !updatedStatus {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(updatedTodo)
	if err != nil {
		return
	}
}

// id is passed as /todos/123
func HandleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deletedTodo, deletedStatus := DeleteTodo(id)
	if !deletedStatus {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(deletedTodo)
}
