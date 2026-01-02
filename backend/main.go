package main

import (
	"fmt"
	"net/http"
)

import (
	"github.com/go-chi/chi/v5"
)

// getUser handles GET requests to fetch a user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User ID: " + id))
}

// createUser handles POST requests to create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}

func main() {
	r := chi.NewRouter()
	r.Use(MyMiddleware)
	r.Get("/users/{id}", getUser)
	r.Post("/users", createUser)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request is received!")
		if r.Method == "GET" {
			fmt.Println("Request is a GET request!")
		} else if r.Method == "POST" {
			fmt.Println("Request is a POST request!")
		} else {
			fmt.Println("Request is a " + r.Method + " request!")
		}
		next.ServeHTTP(w, r)
	})
}
