package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// func validateRequest(req http.Request) error {
//     req.Username = strings.TrimSpace(req.Username)
//     req.State = strings.TrimSpace(req.State)

//     if req.ID == nil || req.Username == "" || req.State == "" {
//         return fmt.Errorf("invalid request")
//     }

//     return nil
// }

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	State    string `json:"state"`
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method mismatch", http.StatusMethodNotAllowed)
	}
	_, _ = w.Write([]byte("Welcome to try /user?name=emeka"))
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if r.Method != http.MethodGet {
		http.Error(w, "Method mismatch", http.StatusMethodNotAllowed)
	}

	res := map[string]any{
		"ok":       true,
		"message":  "Json encoded",
		"dateTime": time.Now().UTC(),
	}
	_ = json.NewEncoder(w).Encode(res) // ecode writes to the json through the encode/json package while decode reads the json

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

func testHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method mismatch", http.StatusMethodNotAllowed)
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "this is a post method",
		})
		return
	}
	defer r.Body.Close()

	var req User

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid json format",
		})
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.State = strings.TrimSpace(req.State)

	if req.ID == 0 || req.State == "" || req.Username == "" {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"success": false,
			"msg":     "Please provide the needed value(s)",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"Success":    true,
		"msg":        req,
		"timestampe": time.Now().UTC(),
	})

}

func main() {

	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/ok", successHandler)

	fmt.Println("Server have started on port :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Server crashed before starting with error: %v \n", err)
	}

}
