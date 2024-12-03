package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func runServer() {
	router := mux.NewRouter()
	router.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		// Create sample student data
		students := []map[string]interface{}{
			{
				"name": "John Doe",
				"age":  20,
				"grades": map[string]interface{}{
					"math":    95,
					"science": 88,
					"history": 92,
				},
			},
			{
				"name": "Jane Smith",
				"age":  21,
				"grades": map[string]interface{}{
					"math":    90,
					"science": 94,
					"history": 89,
				},
			},
		}

		// Set content type header
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode and write JSON response
		json.NewEncoder(w).Encode(students)
	})
	http.ListenAndServe(":8080", router)
}
