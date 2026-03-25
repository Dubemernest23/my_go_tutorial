package main

import (
	"fmt"
	"net/http"
)

func main() {
	// register routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/a", aboutHandler) // registers the function to a pattern more like creating a route

	fmt.Println("Trying to run on port 8080...")

	// start the server — blocks forever (like app.listen() in Express)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error occurred", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Welcome to the home page")
	_, _ = w.Write([]byte("Welcome to go server"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About page")
}
