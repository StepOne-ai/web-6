package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web!")
}

func main() {
	http.HandleFunc("/get", helloHandler)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}