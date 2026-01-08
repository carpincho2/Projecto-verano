package main

import (
	"encoding/json"
	"net/http"
)
[]Products 
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		response := map[string]string{
			"status":  "ok",
			"message": "backend activo",
		}
		json.NewEncoder(w).Encode(response)
	})
	http.ListenAndServe(":8080", nil)
}
