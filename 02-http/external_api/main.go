package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// helper function - write to json
func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// helper function -  fect data
func fectData() (CatFactResponse, error) {
	uri := "https://catfact.ninja/fact"

	resp, err := http.Get(uri)
	if err != nil {
		return CatFactResponse{}, err

	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return CatFactResponse{}, fmt.Errorf("External api failed %v \n", resp.Status)
	}

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return CatFactResponse{}, fmt.Errorf("External api failed: %v", err)
	}

	var data CatFactResponse

	if err := json.Unmarshal(bodyByte, &data); err != nil {
		return CatFactResponse{}, err
	}

	return data, nil

}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	// check the status method
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"success": false,
			"error":   "Get method is not allowed",
		})
		return
	}

	data, err := fectData()
	if err != nil {
		writeJson(w, http.StatusBadGateway, map[string]any{
			"success": false,
			"error":   "Error occurred while fetching an external api",
		})
		return
	}
	writeJson(w, http.StatusOK, map[string]any{
		"success":   true,
		"timestamp": time.Now().UTC(),
		"msg": map[string]any{
			"source": "catfact.ninja",
			"fact":   data.Fact,
			"legth":  data.Length,
		},
	})
}

func main() {

	http.HandleFunc("/external", externalHandler)

	fmt.Println("server have started")

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		fmt.Println("Server crashed due to ", err)
	}
}
