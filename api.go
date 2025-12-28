package main

import (
	"encoding/json"
	"fmt"
	"go-learning/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	data := []map[string]interface{}{{
		"status":  "ok",
		"message": "Server is running",
	},
		{
			"status":  "error",
			"message": "Server is not running",
		},
		{
			"status":  "warning",
			"message": "Server is running with warnings",
		},
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", getHealth).Methods("GET")
	r.HandleFunc("/auth", controllers.GetAuth).Methods("GET")
	fmt.Println("Server is running on port 8080")

	ch := make(chan string)
	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			fmt.Println("Error starting server: ", err)
			return
		}
		ch <- "done"
	}()
	fmt.Println(<-ch)
}
