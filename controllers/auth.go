package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	data := map[string]string{
		"status":  "ok",
		"message": "Auth is running",
	}
	json.NewEncoder(w).Encode(data)
}
