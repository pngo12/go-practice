package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	// Sort of like router, takes in a route, and the function
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	
	http.ListenAndServe(":5000", nil)

}
