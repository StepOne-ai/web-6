package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello, stranger!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/api/user", helloHandler)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}