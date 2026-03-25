package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Server crashed before starting with error: %v \n", err)
	}
}
func baseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method mismatch", http.StatusMethodNotAllowed)
	}
	_, _ = w.Write([]byte("Welcome to try /user?name=emeka"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method mismatch", http.StatusMethodNotAllowed)
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	_, _ = w.Write([]byte(name))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About page")
}
