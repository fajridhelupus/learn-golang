package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("server Starting")
	http.ListenAndServe(":3000", nil)
}
