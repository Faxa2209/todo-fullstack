package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

// Model of Todo item
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	// Create router
	mux := http.NewServeMux()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})

	//Crete array of todos
	todos := []Todo{}

	// Add routes
	mux.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			todo := &Todo{}
			err := json.NewDecoder(r.Body).Decode(todo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			todo.ID = len(todos) + 1
			todos = append(todos, *todo)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todos)
		} else if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todos)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			w.Header().Set("Content-Type", "application/json")

			idParam := r.URL.Path[len("/api/todos/"):]
			id, err := strconv.Atoi(idParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid id"))
				return
			}
			for i, t := range todos {
				if t.ID == id {
					todos[i].Done = !todos[i].Done
					break
				}
			}
			json.NewEncoder(w).Encode(todos)
		case http.MethodDelete:
			w.Header().Set("Content-Type", "application/json")
			idParam := r.URL.Path[len("/api/todos/"):]
			id, err := strconv.Atoi(idParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Invalid id"))
				return
			}
			for i, t := range todos {
				if t.ID == id {
					todos = append(todos[:i], todos[i+1:]...)
					break
				}
			}
			json.NewEncoder(w).Encode(todos)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware.Handler(mux),
	}
	log.Fatal(srv.ListenAndServe())
}
