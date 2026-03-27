package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func useUrL(q http.ResponseWriter, r *http.Request) {
	val := "https://jsonplaceholder.typicode.com/todos"
	resp, err := http.Get(val)
	if err != nil {
		writeJson(q, http.StatusBadRequest, map[string]any{
			"success": false,
			"msg":     "Error occured while recieving data",
		})
		return
	}
	defer resp.Body.Close()

	v := resp.StatusCode
	y := resp.Status
	fmt.Println(y)
	fmt.Println(v)
}

func main() {
	val := "https://jsonplaceholder.typicode.com/todos"
	resp, err := http.Get(val)
	if err != nil {
		fmt.Println(err)
		// writeJson(q, http.StatusBadRequest, map[string]any{
		// 	"success": false,
		// 	"msg":"Error occured while recieving data",
		// })
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return
	}

	v := resp.StatusCode
	fmt.Println(v)
	y := resp.Status
	fmt.Println(y)
	// useUrL()
}
