package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Counter struct {
	Value int
}

var counter Counter

func getCountHandler(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(counter)
}

func updateCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	var update struct {
		Count int `json:"count"`
	}

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if update.Count == 0 {
		http.Error(w, "Это не число", http.StatusBadRequest)
		return
	}

	counter.Value += update.Count
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getCountHandler(w, r)
		} else if r.Method == http.MethodPost {
			updateCountHandler(w, r)
		}
	})

	fmt.Println("Server is listening on port 3333")
	http.ListenAndServe(":3333", nil)
}